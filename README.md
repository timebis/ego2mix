# Ego2mix

This Go library calls the Eco2mix API from RTE and returns data about the French electricity grid.

Those data are real time data.
If you want fully validated data, one shoul refer to this [API](https://odre.opendatasoft.com/explore/dataset/eco2mix-national-cons-def/information/?disjunctive.nature), data are available around one year later.
## deploy
git tag v0.3.2
git push origin v0.3.2
GOPROXY=proxy.golang.org go list -m github.com/timebis/ego2mix@v0.3.2

## Usage
go get github.com/timebis/ego2mix@v0.3.2
see example file