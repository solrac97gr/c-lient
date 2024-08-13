# C-lient

## Description

Make your life easier with C-lient, a command-line tool that helps you manage your C projects.

## Installation

1. **Download the binary** from the [releases page](https://github.com/solrac97gr/c-lient/releases).

2. **Install the binary** using the provided `install.sh` script:
   ```bash
   git clone git@github.com:solrac97gr/c-lient.git
   cd c-lient
   sudo ./install.sh
   cd ..
    ```
3. **Usage**
    ```bash
    c-lient PROJECT_NAME
    ```
    This will create a new directory with the name `PROJECT_NAME` and the following structure:
    ```
    ├── Makefile
    ├── includes
    │   └── utils.h
    ├── src
    │   ├── main.c
    │   └── utils
    │       └── utils.c
    └── tools
        └── info.sh
    ```
    The `Makefile` is a simple Makefile that compiles the project. The `includes` directory contains the `utils.h` header file. The `src` directory contains the `main.c` file and a `utils` directory with the `utils.c` file. The `tools` directory contains a script that prints information about the project.

4. **Compile the project**
    ```bash
    cd PROJECT_NAME
    make
    ```
5. **Run the project**
    ```bash
    ./build/PROJECT_NAME
    ```
6. **Add new entity tool**
    ```bash
    c-lient new-entity ENTITY_NAME
    ```
    This will create a new directory with the name `ENTITY_NAME` and the following structure:
    ```
    |-- includes
    │   └── ENTITY_NAME.h
    |-- src
    │   └── ENTITY_NAME
    │       └── ENTITY_NAME.c
    ```
    The `ENTITY_NAME.c` file contains the implementation of the entity and the `ENTITY_NAME.h` file contains the entity's interface.