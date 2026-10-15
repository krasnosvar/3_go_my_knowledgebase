// Упражнение 07: граница с внешним миром.
//
// Задача: протестируйте CountWords без файлов:
//   - TestCountWords: strings.NewReader("go is  fun\n") -> 3, err == nil;
//   - TestCountWordsReadError: свой тип errReader, чей Read всегда
//     возвращает ошибку; CountWords должна вернуть эту же ошибку (errors.Is).
//
// Идея: зависимость от io.Reader — это «шов», через который unit-тест
// подменяет файл/сеть. Настоящий файл — уже интеграционный тест (тема 06).
//
// Запуск: go test ./03_exercises/07_integration_boundary/
package words

import "testing"

func TestCountWords(t *testing.T) {
	t.Fatal("TODO: проверьте подсчёт слов и проброс ошибки чтения")
}
