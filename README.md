 ![](https://github.com/viniciusbds/navio/workflows/build/badge.svg)  ![](https://github.com/viniciusbds/navio/workflows/unit-tests/badge.svg) [![godocs](https://godoc.org/github.com/viniciusbds/navio?status.svg)](https://godoc.org/github.com/viniciusbds/navio) 

 
# Navio

<img src="/cargueiro.png" alt="drawing" width="120"/>

----------------------------

**Navio** is a simple tool to create linux containers based on the namespace and cgroups features. 

The Navio creates containers, that is, **a set of processes isolated by Linux namespaces**, for example: PID to isolate the processes and Mount to isolate the file systems.

All created containers have their own **rootfs** (a mini operating system) associated, so that a change (for example, an installation of any library) in a container does not affect others ones.

It is also possible to limit the amount of resources that each container can use, this is done through Cgroups.


### Why?
Just for science, do not use this code in production !!! :D.

## Available Images

| Image| version| size |
| ---- | -----| ------|
| alpine|  v3.11| 2.7M|
| busybox| v4.0| 1.5M|
| ubuntu| v20.04| 90M|

## [Namespaces](https://en.wikipedia.org/wiki/Linux_namespaces)

what the processes can see

**CLONE_NEWUTS** : The UTS namespace provides isolation of the hostname and domainname system identifiers

**CLONE_NEWPID** : PID namespace isolates the process ID number space. This means that two processes running on the same host can have the same PID!

**CLONE_NEWNS** : The Mount namespace isolate the filesystem mount points

---


- [x] UTS - isolate **hostname and domainname**

- [x] PID - isolate the **PID number space**

- [x] MNT - isolate **filesystem mount points**

- [ ] IPC - isolate **interprocess communication (IPC)** resources

- [ ] NET - isolate **network interfaces**

- [ ] User - isolate **UID/GID number spaces**

- [ ] Cgroup - isolate **cgroup root directory**

- [ ] Time Namespace - allows processes to see **different system times** in a way similar to the UTS namespace.


## [Cgroups](https://en.wikipedia.org/wiki/Cgroups)

what the processes can use

- [ ] Memory

- [ ] CPU

- [ ] I/O

- [ ] Process numbers


## Requirements

- [golang environment](https://golang.org/)
- make
- wget
- mysql
- some of commands (ex.: `navio build`, `navio run`, `navio rmi` and `navio exec`) must be executed with sudo privilegies.

## How to install

#### If you just want use, is very simples: 

```
 git clone https://github.com/viniciusbds/navio.git
 cd navio
 ./install.sh
```

#### If you want compile the code before install:

```
 git clone https://github.com/viniciusbds/navio.git
 cd navio
 make
 ./install.sh
```

#### To run all unit tests, type:

```
 cd /path/to/project/navio
 sudo make unit-tests
```

#### To uninstall:

 ```
 cd navio
 ./uninstall.sh
```
  
## Example Commands

`navio get images`

`sudo navio rmi alpine`

`navio pull alpine`

`sudo navio run alpine sh --name mycontainer`

`sudo navio exec mycontainer sh` 

...

`sudo navio run busybox sh`

`sudo navio run ubuntu /bin/bash --name python3apps`

  
## Contributing

You can contribute to the project in any way you want, either by fixing bugs, implementing new features, improving the documentation or proposing new features through issues

See [Contributting](/CONTRIBUTING.md) for more details

## References

  - [Containers from Scratch • Liz Rice](https://www.youtube.com/watch?v=8fi7uSYlOdc)
  
  - [Containers from Scratch](https://ericchiang.github.io/post/containers-from-scratch/)
  
  - [Building a container with less than 100 lines in Go](https://www.infoq.com/br/articles/build-a-container-golang/)

  - [Linux Namespaces](https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a)
  
  - [Namespaces](https://escotilhalivre.wordpress.com/2015/08/12/namespaces/)
  
  - <div><a href="/cargueiro.png" title="Icon">Icon</a> made by <a href="https://www.flaticon.com/br/autores/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/br/" title="Flaticon">www.flaticon.com</a></div>
