
# TF Compose 👩‍🎤
 
TF-compose is a stack composer based on 
[Web Modules](https://INSERTLINK).

This binary application generates a fully functionnal Terraform stack from a YAML description.

You only have to run the generation and deploy your stack, as is from local development to the production servers.

## Quickstart

### Step 1: install tf-compose
Install TF compose:
// @todo storage sur github de chaque release ?
```bash
wget https://tf-compose.matthieuwerner.com/download/tf-compose -O tf-compose
sudo chmod a+x tf-compose
sudo mv tf-compose /usr/local/bin/tf-compose
```

### Step 2: create configuration

Create a deployment folder in your application, we are using "deploy"
in our examples, but you can choose any name you like.
```bash
mkdir [your project root]/deploy
```

Add a valid configuration file to the root of your project.
```yaml
# Default workspace, local environment
default:
    provider: docker
    domain: tf-compose.local
    modules: &modules
      function-php:
        name: A PHP function
        application-path: ../application
        document-root: public
        handler: public/index.php
        php-version: 8.1
        subdomain: www

# Prod workspace, a cloud environment
prod:
    provider: aws
    domain: tf-compose.com
    modules: *modules
```
Read more abour configuration [here](./doc/configuration.md).

### Step 3: Generate files

#### Option 1: generate entire terraform environment

```bash
cd [your project root]/deploy
tf-compose init
tf-compose install
```

#### Option 2: add tf-compose in existing terraform environment

```bash
cd [your project root]/deploy
tf-compose init
tf-compose install
```


Initialize a new Terraform project in your application. You can start with the following configuration:
```terraform
terraform {
  required_version = "> 1.0"

  // In this example, we are using 2 providers, Docker, for local stack, and AWS for cloud stack
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.48.0"
    }
  }

  // Local backend
  backend "local" {}
}

# TF compose
module "tf-compose" {
  source = "./web-modules/tf-compose"
}
```

## More about TF Compose

tf copmpose sert à génrer un fichier tf par workspace avec leßmodules confoguirés dans le yaml
lié aux modules blablabla

- [Documentation](doc/documentation.md)
- [Configuration](doc/configuration.md)

// Genere un grand fichier avec les confs de tous els workspaces dedans à inclure dans son projet (export path)
// Features
// - telecharge ou non les modules (configuration)
// - lit un yaml de configuration de modules et générique aussi comme composer en fait
// - config en json ou yaml
// - gestion des versions avec des mots clés de version ?
// - genere un fichier de loading de modules (par workspace, defautl defautl sinon autre)
//