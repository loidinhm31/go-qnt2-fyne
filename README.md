Packaging application

````
https://developer.fyne.io/started/packaging
````

````shell
go install fyne.io/fyne/v2/cmd/fyne@latest
````

````shell
fyne package -appVersion 1.0.0 -appID QNM1.0 -name QN2Management -release
````

Bundle File
````shell
fyne bundle unreachable.png >> unreachableBundle.go
````