package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VMImageResource{}
var _ resource.ResourceWithImportState = &VMImageResource{}

func NewVMImageResource() resource.Resource {
	return &VMImageResource{}
}

// VMImageResource defines the resource implementation.
type VMImageResource struct {
	client *CudoClientData
}

// VMImageResourceModel describes the resource data model.
type VMImageResourceModel struct {
	ID           types.String                `tfsdk:"id"`
	DataCenterId types.String                `tfsdk:"data_center_id"`
	SizeGib      types.Int64                 `tfsdk:"size_gib"`
	Source       *VMImageSourceResourceModel `tfsdk:"source"`
}

type VMImageSourceResourceModel struct {
	// SnapshotID types.String `tfsdk:"snapshot_id"`
	VmID types.String `tfsdk:"vm_id"`
}

func (r *VMImageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_vm_image"
}

func (r *VMImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Image resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Image ID",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the data center where the image is located.",
				Computed:            true,
			},
			"size_gib": schema.Int64Attribute{
				MarkdownDescription: "Size of the image in GiB",
				Computed:            true,
			},
			"source": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"vm_id": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						MarkdownDescription: "The ID of the VM with the disk to generate an image from.",
						Required:            true,
					},
				},
				Required:            true,
				MarkdownDescription: "The source vm disk or snapshot to generate the image from",
			},
		},
	}
}

func (r *VMImageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *CudoClientData got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *VMImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *VMImageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createParams := virtual_machines.NewCreatePrivateVMImageParamsWithContext(ctx)
	createParams.ProjectID = r.client.DefaultProjectID
	createParams.ID = state.ID.ValueString()
	if state.Source == nil {
		resp.Diagnostics.AddError(
			"Unable to create image resource",
			"Source required to create image resource",
		)
		return
	}
	createParams.VMID = state.Source.VmID.ValueString()
	// createParams.SnapshotID = state.Source.SnapshotID.ValueStringPointer()
	res, err := r.client.Client.VirtualMachines.CreatePrivateVMImage(createParams)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create image resource",
			err.Error(),
		)
		return
	}

	state.DataCenterId = types.StringValue(res.Payload.Image.DataCenterID)
	state.SizeGib = types.Int64Value(int64(res.Payload.Image.SizeGib))

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMImageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *VMImageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	getParams := virtual_machines.NewGetPrivateVMImageParamsWithContext(ctx)
	getParams.ID = state.ID.ValueString()
	getParams.ProjectID = r.client.DefaultProjectID

	res, err := r.client.Client.VirtualMachines.GetPrivateVMImage(getParams)
	if err != nil {
		if apiErr, ok := err.(*virtual_machines.GetPrivateVMImageDefault); ok && apiErr.IsCode(404) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read image resource",
			err.Error(),
		)
		return
	}

	state.ID = types.StringValue(res.Payload.Image.ID)
	state.DataCenterId = types.StringValue(res.Payload.Image.DataCenterID)
	state.SizeGib = types.Int64Value(int64(res.Payload.Image.SizeGib))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VMImageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error getting vm image plan",
			"Error getting vm image plan",
		)
		return
	}

	// Read Terraform state data into the model
	var state NetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *VMImageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *VMImageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	params := virtual_machines.NewDeletePrivateVMImageParamsWithContext(ctx)
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ID.ValueString()

	_, err := r.client.Client.VirtualMachines.DeletePrivateVMImage(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete image resource",
			err.Error(),
		)
		return
	}
}

func (r *VMImageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
