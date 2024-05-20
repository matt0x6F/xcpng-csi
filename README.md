# XCP-ng CSI Driver

XCP-ng CSI Driver is a Go project that uses the Container Storage Interface to integrate native XCP-ng Volumes in Kubernetes or other CSI capable container orchestrators. This driver communicates with your Xen Cluster via Xen Orchestra appliance.

This software has no tests and is largely filling a need in a hobby project. Do with that information what you will.

## Installation

Use the package manager Helm to install the driver in Kubernetes

> MAKE SURE that you update `values.yaml` according to your config 

```bash
helm upgrade --install --namespace=xcpng-csi --create-namespace xcpng-csi ./chart -f ./values.yaml
```

## Usage

Create StorageClass that consumes driver

```yaml
# StorageClass that uses the XCP-ng CSI Driver
apiVesrion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: fast
provisioner: csi.xcpng.ooo-yay.com
parameters:
   Datastore: "Optional: Storage Repository Name"
   FSType: "Optional: Filesystem Type ie. ext4 Defaults to ext4"
```

Then create PersistentVolumeClaim that uses StorageClass

```yaml
# PersistentVolumeClaim that users the fast StorageClass
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
   name: test-claim
spec:
   accessModes:
     - ReadWriteOnce
   resources:
     requests:
       storage: 5Gi
   storageClassName: fast
---
# Pod that consumes the PersistentVolumeClaim
apiVersion: v1
kind: Pod
metadata:
   name: test-pod
spec:
   volumes:
     - name: config
       persistentVolumeClaim:
         claimName: test-claim
   containers:
     - name: test-pod
       image: nginx:latest
       volumeMounts:
         - name: config
           mountPath: /usr/share/nginx/html
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Credit

- [Original work](https://github.com/ArturoGuerra/xcpng-csi) by Arturo Guerra
- [Updates](https://github.com/shokre/xcpng-csi) by @shokre

## License
[MIT](https://choosealicense.com/licenses/mit/)
