// Desired schema, read from the gorm models via tools/atlasloader.
// The loader lives in its own module (tools/go.mod) so the Atlas provider's
// dependencies stay out of the bot's build, hence -C tools.
data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-C",
    "tools",
    "./atlasloader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url

  // Throwaway database Atlas uses to compute diffs. Start it with:
  //
  //   docker run --rm -d --name atlas-dev -p 3307:3306 \
  //     -e MYSQL_ROOT_PASSWORD=atlas -e MYSQL_DATABASE=dev mysql:9.2.0
  //
  // Atlas's own docker:// shorthand is not usable here: it dials the container
  // on a Docker Desktop internal address that the Windows host cannot reach.
  // Running it ourselves also means matching production's MySQL version, which
  // the arigaio/* images do not publish.
  //
  // Atlas wipes this schema on every run, so it must never point at
  // phoenix_bot_db. Port 3307 keeps it clear of the compose database.
  dev = "mysql://root:atlas@localhost:3307/dev"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
