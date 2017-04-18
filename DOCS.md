Use this plugin for caching build artifacts to speed up your build times. This
plugin can create and restore caches of any folders.

## Config

The following parameters are used to configure the plugin:

* **host** - full URI of sonar host - sonar.host.url property
* **login** - authenticate with this user against sonar server - sonar.login property
* **password** - authenticate with this password against sonar server - sonar.password property
* **key** - Sonar project key. Defaults to ORG:REPO_NAME - sonar.projectKey property 
* **name** - Sonar project name. Defaults to ORG/REPO_NAME - sonar.projectName property
* **version** - Sonar project version - sonar.projectName property
* **sources** - Sonar project sources paths - sonar.sources property
* **inclusions** - Sonar project sources inclusions - sonar.inclusions property
* **exclusions** - Sonar project sources exclusions - sonar.exclusions property
* **language** - Sonar project language. Defaults to 'js' - sonar.language property
* **profile** - Sonar project profile. Defaults to 'node' - sonar.profile property
* **encoding** - Sonar project encoding. Defaults to 'UTF-8' - sonar.sourceEncoding property
* **lcovpath** - Sonar project lcov coverage file. Defaults to 'test/coverage/reports/lcov.info' - sonar.javascript.lcov.reportPath property


The following secret values can be set to configure the plugin.

* **SONAR_HOST** - corresponds to **host**
* **SONAR_LOGIN** - corresponds to **login**
* **SONAR_PASSWORD** - corresponds to **password**

It is highly recommended to put the **SONAR_LOGIN** and
**SONAR_PASSWORD** into a secret so it is
not exposed to users. This can be done using the drone-cli.

```bash
drone org secret add THEORG SONAR_LOGIN theuser --image ypcloud/sonar-runner
```

```bash
drone org secret add THEORG SONAR_PASSWORD thepassword --image ypcloud/sonar-runner
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:

  build:
    image: node:latest
    commands:
      - npm install
      - npm test

  sonar:
    image: ypcloud/sonar-runner
    sources: "service/,utils/,daemon/,ui/"
    version: "0.1.0"    
```

