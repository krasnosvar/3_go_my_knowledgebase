# Kubernetes Operators

Operator — приложение, расширяющее Kubernetes собственным контроллером и Custom Resource Definition (CRD). Управляет жизненным циклом сложных приложений (БД, очереди) внутри кластера.

## Когда нужен

- Сложное приложение требует многошагового management (БД-кластер, scaling, failover, backup).
- Вы хотите kubernetes-native управление — `kubectl get mydb`.
- Команда уже на Kubernetes и готова поддерживать Go-codebase.

## Когда НЕ нужен

- Простое приложение → достаточно Deployment/StatefulSet/Helm.
- Нет реальной автоматизации (operator делает то, что и так делает Helm).
- Команда не имеет k8s-опыта — operator под нагрузкой требует серьёзной экспертизы.

## Стек

- **kubebuilder** — генератор и фреймворк (стандарт).
- **operator-sdk** — альтернатива, но kubebuilder сейчас доминирует.
- **controller-runtime** — runtime под контроллер, общий для обоих.

## Concepts

- **Custom Resource (CR)** — объект с собственной схемой, например `MyDatabase`.
- **CRD** — определение схемы CR.
- **Controller** — реконсиляционный цикл, наблюдающий CR'ы и приводящий реальное состояние к желаемому.

## Reconcile loop

Сердце controller'а:

```go
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var db myv1.MyDatabase
    if err := r.Get(ctx, req.NamespacedName, &db); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // 1. Проверить, что желаемое = реальное
    // 2. Если нет — создать/обновить нужные resources
    // 3. Обновить status
    // 4. Решить, нужен ли requeue

    return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}
```

`Reconcile` вызывается на каждое изменение CR (или watched ресурса).

## Idempotency

Reconcile может вызваться 100 раз. Каждый вызов должен:

- Не делать лишнего, если состояние уже нужное.
- Использовать `CreateOrUpdate` патчи.
- Обновлять status осмысленно.

## Status и Conditions

Status описывает текущее состояние:

```go
type MyDatabaseStatus struct {
    Phase      string             `json:"phase,omitempty"`
    Conditions []metav1.Condition `json:"conditions,omitempty"`
    Replicas   int32              `json:"replicas"`
}
```

Conditions — массив с `Type`/`Status`/`Reason`/`Message`. Стандартный k8s-паттерн для отображения «Ready», «Progressing», «Degraded».

## Finalizers

Если operator управляет внешними ресурсами (S3 bucket, DNS-запись), при удалении CR нужно их вычистить. Finalizer блокирует удаление, пока operator не завершит cleanup:

```go
if obj.DeletionTimestamp != nil {
    if controllerutil.ContainsFinalizer(&obj, finalizerName) {
        if err := r.cleanupExternal(ctx, &obj); err != nil { return ... }
        controllerutil.RemoveFinalizer(&obj, finalizerName)
        if err := r.Update(ctx, &obj); err != nil { return ... }
    }
    return ctrl.Result{}, nil
}
```

## Тестирование

- Unit на reconcile-функции с fake client.
- Integration — на envtest (локальный API-server без полного k8s).
- E2E — на kind / k3d с реальными CRD.

## Правило

- Operator — серьёзная инвестиция. Сначала Deployment/Helm, и только если автоматизация реально нужна — operator.
- Reconcile идемпотентен.
- Finalizers для всего внешнего state.
- Status с conditions — для видимости.
