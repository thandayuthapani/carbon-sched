# Meeting Agenda and Notes
## July 8, 2022
### Agenda
- Discussion on metrics sources - [WattTime](https://www.watttime.org/api-documentation/#introduction), [ElectricityMap](https://static.electricitymap.org/api/docs/index.html)

### Notes
- Alternate scheduling schemes for comparision

## July 21, 2022
### Agenda

- ElectricityMap data - [2021_Germany](https://drive.google.com/file/d/1F-PXu4p9sR28Gx2aJII-kuesCWV5vPS9/view?usp=sharing), [metrics](https://docs.google.com/spreadsheets/d/e/2PACX-1vQymR9eNK7U9bDSUBlyegx0y6FPhpe-mVBGniPzGtWDjZyHb8gI2NHSx-S49EXBhCkDe8dqfJAvsi3C/pubhtml#)
- Discussion on admiralty architecture - [link](https://admiralty.io/docs/concepts/topologies)
- Discussion on type of functions
- Metrics comparision - Latency, Carbon Efficiency
- Alternate Scheduling schemes - default, carbon-sched, geo-location sched
- Absolute Metrics? Carbon-efficient region vs deployment location

### Notes
- Deploying Admiralty on cluster
- Read literature on metrics - [1](https://dl.acm.org/doi/10.1145/3530688), etc

## August 4th, 2022
### Agenda

- Update on metrics collector
- Update on deploying Admiralty on cluster
- Power Usage Effectiveness (PUE) - A plausible choice?

### Notes
- Storing carbon metrics - Is it necessary? Can we use historical API to collect the history data?
- Debug admiralty deployment in cluster

## August 18th, 2022
### Agenda
- Update on metrics collector - Cluster ready
- Update on admiralty deployment - Tested in cluster with mock deployment of pods
- Discussion on [Cucumber](https://link.springer.com/chapter/10.1007/978-3-031-12597-3_14) literature

### Notes
- Add normalization code in metrics-collector
- Add scheduling plugin code to do carbon-aware scheduling
- Integrate Knative into the setup

## September 5th, 2022
### Agenda
- Update on metrics-collector and scheduling plugin

### Notes
- Scheduling Plugin
- latency in admiralty

## September 19th, 2022
### Agenda
- East-West communication issue in admiralty - [admiralty](https://admiralty.io/pricing/), [Cilium](https://cilium.io/blog/2019/03/12/clustermesh/)
- Alternative for admiralty - [multi-cluster mesh](https://www.cncf.io/blog/2021/04/12/simplifying-multi-clusters-in-kubernetes/), [liqo](https://github.com/liqotech/liqo)
- Update on scheduler testing in cluster

### Notes
- Evaluate Liqo for multi-cluster scheduling

## September 27th
### Agenda
- Update on evaluation of liqo to our usecase

### Notes
- Hack22 hackathon

## October 10th
### Agenda
- Update on scheduler code
- Update on Carbon SDK - how to use carbon sdk in our hackathon usecase
- Discussion on hackathon

### Notes
- Finish up on scheduler code

## October 20, 2022
### Agenda
- Update on scheduler code
- Update on metrics-server

### Notes
- Finish scheduler code
- Make changes to metrics-collector to collect data from carbon-sdk
- Load tester for functions
- Upload functions to the repo

## October 27, 2022
### Agenda
- Update on scheduler code
- Update on metrics-collector

### Notes