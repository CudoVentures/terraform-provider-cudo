# start a vm in any data center by specifying a maximum price
resource "cudo_vm" "my-vm-max-price" {
  id         = "terra-vm-1"
  memory_gib = 2
  vcpus      = 1
  boot_disk = {
    image_id = "debian-11"
  }
  max_price_hr = 0.003
  networks = [
    {
      network_id         = "my-network"
      assign_public_ip   = true
      security_group_ids = ["my-security-group"]
    }
  ]
}

# pick a specific data center and machine type
resource "cudo_vm" "my-vm" {
  id             = "terra-vm-1"
  machine_type   = "standard"
  data_center_id = "gb-bournemouth-1"
  memory_gib     = 2
  vcpus          = 1
  boot_disk = {
    image_id = "debian-11"
    size_gib = 50
  }
  ssh_key_source = "project"
  start_script   = <<EOF
                     touch /multiline-script.txt
                     echo  $PWD > /current-dir.txt
                     EOF
}
