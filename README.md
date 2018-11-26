# alplab-client-v2
This client is a command line tool for developers to use services of ALP.Lab individually.

## Examples

#### Road Side Sensor to Simulation Transformation  

This transforms roadside sensor data from ASFINAG flow of traffic services 
on selected Austrian Roads into the ego car's coordinate system. 
The transformation is part of ALP.Lab cloud services; the results are 
returned into the [Open Simulation Interface](https://github.com/OpenSimulationInterface/open-simulation-interface)
to be of direct use in your simulation applications. 

In these examples the files with roadside and car sensor data as well as the config file were placed in the binaries' directory itself, but they can also be some place else and be referenced relative or absolute with the corresponding flag.

Bash Command Example:

```
./alp rtrans -r radar.json -c car.gpx
```

Windows Command Example:

```
alp.exe -r radar.json -c car.gpx
```

The config parser can read from [YAML](http://yaml.org/), [JSON](https://json.org/), [TOML](https://github.com/toml-lang/toml), [HCL](https://github.com/hashicorp/hcl), and [Java properties](https://docs.oracle.com/javase/7/docs/api/java/util/Properties.html) config files.

alp.yml Example:

```
host: radar-transform.XXX.xx
port: 443
certificate: ca-chain.cert.pem
```
