// @tg version=1.0.0
// @tg title=`Example API`
// @tg description=`Simple summator`
// @tg servers=`http://localhost:9000`
//go:generate tg client -go --services . --outPath ../clients/adder
//go:generate tg transport --services . --out ../transport --outSwagger ../../api/adder-openapi.yaml
//go:generate goimports -l -w ../transport ../clients/adder
package interfaces
