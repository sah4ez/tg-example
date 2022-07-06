// @tg version=1.0.0
// @tg title=`Example API`
// @tg description=`Simple summator`
// @tg servers=`http://localhost:9000`
//go:generate tg client -go --services . --outPath ../clients/user
//go:generate tg client -js --services . --outPath ../../js
//go:generate tg transport --services . --out ../transport --outSwagger ../../api/user-openapi.yaml
//go:generate goimports -l -w ../transport ../clients/user
package interfaces
