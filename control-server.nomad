job "drone-control" {
  datacenters = ["dc1"]

  type = "service"

  constraint {
    attribute = "${attr.cpu.arch}"
    value     = "arm"
  }

  update {
    max_parallel = 1
    min_healthy_time = "10s"
    healthy_deadline = "3m"
    auto_revert = false
    canary = 0
  }

  group "drone-control" {
    count = 1

    restart {
      # The number of attempts to run the job within the specified interval.
      attempts = 10
      interval = "5m"

      delay = "25s"

      mode = "delay"
    }

    ephemeral_disk {
      size = 10
    }

    task "webserver" {
      # The "driver" parameter specifies the task driver that should be used to
      # run the task.
      driver = "exec"

      # The "config" stanza specifies the driver configuration, which is passed
      # directly to the driver to start the task. The details of configurations
      # are specific to each driver, so please see specific driver
      # documentation for more information.
      config {
        command = "control"
        args = [
          "Mambo_586378"
        ]
      }

      artifact {
        source = "./control"

      }

      logs {
        max_files     = 2
        max_file_size = 3
      }

      resources {
        cpu    = 100 # 500 MHz
        memory = 128 # 256MB
        network {
          mbits = 10
        }
      }
    }
  }
}
