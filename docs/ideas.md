# Ideas, mejoras, bugs, etc

## Mejoras

- [ ] `cleaner.go` Actualmente admite dos parámetros: `-loop` y `-interval`, pero en realidad siempre se tienen que proporcionar los dos parámetros de forma conjunta. Se puede mejorar *fusionando* los dos parámetros en uno solo, que indique se de tiene que ejecutar en forma de *loop* y qye admita como argumento cada cuánto se tiene que ejecutar el *lopp*. Básicamente, sería *renombrar* `-interval` a `-loop` y modificar el comportamiento para que, si está presente, se ejecute en bucle con el tiempo especificado.

  Tal y como está el código ahora, habría que modificar el valor por defecto de `-interval`/`-loop` para que fuera `0s` y que en ese caso no se entre en el bucle *for*.
