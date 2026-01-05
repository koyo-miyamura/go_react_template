env "local" {
  src = "ent://internal/infra/ent/schema"
  dev = "mysql://root@db:3306/dev"
  url = "mysql://root@db:3306/template_db"

  migration {
    dir = "file://internal/infra/ent/migrate/migrations"
  }

  diff {
    skip {
      drop_schema = true
      drop_table  = true
    }
  }
}
