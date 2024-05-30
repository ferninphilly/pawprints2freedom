remote_state {
  backend = "s3"
  config = {
    encrypt = true
    bucket  = "${basename(path_relative_to_include())}-terraform-pptf-service-tfstate"
    key     = "stacks/${basename(path_relative_to_include())}/lambdas/terraform.tfstate"
    region  = "eu-west-1"
    profile = "pawprints2freedom"
    dynamodb_table = "terraform-state-lock-dynamo-${basename(path_relative_to_include())}"
  }
}

terraform {
  /* Can use hooks to run arbitrary shell commands, for instance for compiling binaries */
  before_hook "compile_lambda" {
    commands     = ["plan"]
    execute      = ["${get_parent_terragrunt_dir()}/build_scripts/compile.sh", "${get_parent_terragrunt_dir()}/infrastructure_module/lambdas/code"]
    run_on_error = false
  }
  /* plan args */
  extra_arguments "plan_args" {
    commands = [
      "plan"
    ]

    arguments = [
      "-out=${get_terragrunt_dir()}/terraform_${basename(path_relative_to_include())}.tfplan",
      "-lock=true",
      "-lock-timeout=60s"
    ]
  }

  //apply args
    extra_arguments "apply_args" {
    commands = [
        "apply"
    ]

    arguments = [
      "-lock=true",
      "-lock-timeout=60s",
      "${get_terragrunt_dir()}/terraform_${basename(path_relative_to_include())}.tfplan"
    ]
}

  extra_arguments "vars" {
    commands = [
      "plan",
      "import",
      "push",
      "refresh",
      "destroy"
    ]

    # .tfvars for set variables
    # .vars for declared variables
    // required_var_files = [
    //   "${get_terragrunt_dir()}/../../${basename(path_relative_to_include())}.tfvars"
    // ]

    # optional_var_files = [
    # ]

  }
}

inputs = {
  environment = "${basename(path_relative_to_include())}"
  region      = "eu-west-1"
  account_id  = "637423452555"
}
