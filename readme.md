# dcli-other-projects: Sample Projects for dcli

## Overview

This repository contains example project setups for use with the [dcli](https://github.com/mh-daneshvar/dcli) tool. It provides sample Docker Compose configurations to demonstrate how you can define and manage multiple projects and services using `dcli`.

The structure of each project and service follows a typical setup with separate Docker Compose files for each service. These samples serve as a starting point or reference when configuring your own projects with `dcli`.

## Repository Structure

Each project directory contains its services, and each service has its own `docker-compose.yaml` file. You can customize these files to suit your actual project needs.

### Example Folder Structure:

```bash
dcli-other-projects/
├── project-one/
│   ├── auth-service/
│   │   └── docker-compose.yaml
│   ├── users-service/
│   │   └── docker-compose.yaml
│   ├── catalogue-service/
│   │   └── docker-compose.yaml
│   └── common-services/
│       └── docker-compose.yaml
└── project-two/
    ├── order-service/
    │   └── docker-compose.yaml
    ├── payment-service/
    │   └── docker-compose.yaml
    └── common-services/
        └── docker-compose.yaml
```

### Explanation of Folder Structure:

- **project-one/** and **project-two/**: Each directory represents a project with several services.
- **auth-service/**, **users-service/**, **catalogue-service/**, etc.: These directories contain individual services, each with its own Docker Compose file (`docker-compose.yaml`).
- **common-services/**: Shared services between different project services, like logging or web UIs, managed through their own Docker Compose file.

## How to Use

1. **Clone this repository**:

   ```bash
   git clone https://github.com/mh-daneshvar/dcli-other-projects.git
   ```

2. **Configure with dcli**:

   Use the `dcli` tool to manage these projects and services. Point the `docker_compose_file_path` in your `dcli` configuration (YAML) to the appropriate service paths within this repository.

3. **Run dcli**:

   After setting up your configuration, run the `dcli` CLI to select and manage projects interactively.

## Example

If you are using `dcli` to manage `project-one`, the `docker_compose_file_path` for the **Auth Service** in your `dcli` YAML configuration might look like this:

```yaml
docker_compose_file_path: ./dcli-other-projects/project-one/auth-service/docker-compose.yaml
```

This file structure and paths will allow `dcli` to correctly manage the Docker containers defined in this repository.

## Contributing

If you'd like to contribute additional projects or services:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/YourFeatureName`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to your branch (`git push origin feature/YourFeatureName`).
5. Open a pull request.

## License

This repository is distributed under the MIT License. See the `LICENSE` file for more details.