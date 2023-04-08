variable "hello" {
  type = string
}

variable "stars" {
  type = number
}

variable "cars" {
  type = number
}

output "hello" {
  value = var.hello
}

output "stars" {
  value = var.stars
}

output "cars" {
  value = var.cars
}

output "secret" {
  value     = "super-secret!"
  sensitive = true
}

output "happy" {
  value = true
}
