# path/filepath

Пакет `filepath` — работа с путями, учитывающая ОС (слэши на Unix, обратные слэши на Windows).

## path vs path/filepath

- `path` — для URL-подобных путей и slice-путей с `/` всегда.
- `path/filepath` — для файловой системы. Используй его в прикладном коде.

## Базовые операции

```go
filepath.Join("a", "b", "c")       // "a/b/c" (unix) или "a\\b\\c" (windows)
filepath.Split("/a/b/file.txt")    // "/a/b/", "file.txt"
filepath.Dir("/a/b/file.txt")      // "/a/b"
filepath.Base("/a/b/file.txt")     // "file.txt"
filepath.Ext("file.tar.gz")        // ".gz"
filepath.Abs("relative/path")       // абсолютный путь
filepath.Rel(base, target)          // относительный путь
filepath.Clean("/a//b/../c")        // "/a/c"
```

## Walk

Обход директорий:

```go
err := filepath.WalkDir("root", func(path string, d fs.DirEntry, err error) error {
    if err != nil { return err }
    if d.IsDir() && d.Name() == ".git" {
        return filepath.SkipDir
    }
    fmt.Println(path)
    return nil
})
```

`WalkDir` быстрее `Walk`: не вызывает stat на каждом файле.

## Match / Glob

```go
ok, _ := filepath.Match("*.go", "main.go")       // true
matches, _ := filepath.Glob("*.go")              // все .go в текущей
```

Простые wildcard: `*`, `?`, `[abc]`, `[a-z]`. Не поддерживает `**` — для рекурсии используй `WalkDir`.

## Symlinks

```go
target, err := filepath.EvalSymlinks(p) // резолвит все симлинки
```

`filepath.Abs` их не резолвит — только делает путь абсолютным.

## Кроссплатформенность

Всегда используй `filepath.Join` вместо ручной склейки через `/`. Даже если кажется, что код работает только на Linux — лучше иметь одну привычку.

## io/fs vs os

С Go 1.16+ есть пакет `io/fs`, абстракция над файловой системой. Позволяет тестировать код, подставляя in-memory FS.

```go
import "io/fs"

func listTxt(fsys fs.FS) []string {
    var out []string
    _ = fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
        if filepath.Ext(p) == ".txt" { out = append(out, p) }
        return nil
    })
    return out
}
```

В проде `os.DirFS("/data")`, в тестах `fstest.MapFS{...}`.

## Частые ошибки

- Ручная склейка `dir + "/" + file` — на Windows ломается.
- Путать `path.Join` (URL) и `filepath.Join` (FS).
- Не проверять `filepath.Rel` на ошибку — может вернуть сумасшедший путь, если базы не пересекаются.
- Забыть, что `Match` поддерживает только простые паттерны. Для сложных — regexp или свой парсер.
