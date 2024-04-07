# Design pattern

## 1. Singleton
El patrón Singleton asegura que una clase tenga solo una instancia y proporciona un punto de acceso global a esa instancia. Es útil cuando se necesita un control estricto sobre cómo y cuándo se accede a ciertos recursos compartidos.
## 2. Observador
El patrón Observador define una dependencia de uno-a-muchos entre objetos de manera que cuando un objeto cambia su estado, todos sus dependientes son notificados automáticamente. Es ampliamente utilizado para implementar sistemas de eventos.
## 3. Fábrica (Factory)
El patrón Fábrica define una interfaz para crear un objeto, pero deja que las subclases decidan qué clase instanciar. Facilita la adición de nuevas variantes de productos sin alterar el código que usa la fábrica.
## 4. Abstract Factory
El patrón Abstract Factory proporciona una interfaz para crear familias de objetos relacionados o dependientes sin especificar sus clases concretas. Es útil cuando los sistemas deben ser independientes de cómo se crean, componen y representan sus productos.
## 5. Builder
El patrón Builder separa la construcción de un objeto complejo de su representación, de modo que el mismo proceso de construcción puede crear diferentes representaciones. Es útil cuando se debe construir un objeto con muchas opciones posibles.
## 6. Prototipo (Prototype)
El patrón Prototipo crea nuevos objetos clonándolos de un objeto existente, facilitando la adición de cualquier subclase de un objeto conocido en tiempo de ejecución.
## 7. Adaptador (Adapter)
El patrón Adaptador permite que interfaces incompatibles trabajen juntas. Actúa como un intermediario entre dos clases, convirtiendo la interfaz de una clase en otra que el cliente espera.
## 8. Decorador (Decorator)
El patrón Decorador añade responsabilidades adicionales a un objeto de manera dinámica. Los decoradores proporcionan una alternativa flexible a la subclase para extender la funcionalidad.
## 9. Fachada (Facade)
El patrón Fachada proporciona una interfaz unificada a un conjunto de interfaces en un subsistema. Define una interfaz de nivel más alto que hace que el subsistema sea más fácil de usar.
## 10. Puente (Bridge)
El patrón Puente desacopla una abstracción de su implementación, de modo que las dos puedan variar de forma independiente. Es útil cuando se quiere evitar un enlace permanente entre la abstracción y su implementación.
## 11. Comando (Command)
El patrón Comando convierte una solicitud en un objeto independiente que contiene toda la información sobre la solicitud. Esto permite parametrizar métodos con diferentes solicitudes, retrasar o poner en cola la ejecución de una solicitud, y soportar operaciones que pueden deshacerse.
## 12. Cadena de Responsabilidad (Chain of Responsibility)
Este patrón pasa la solicitud a lo largo de una cadena de manejadores potenciales hasta que uno de ellos maneja la solicitud. Es útil para desacoplar a los remitentes y receptores de una solicitud.
## 13. Estrategia (Strategy)
El patrón Estrategia define una familia de algoritmos, encapsula cada uno de ellos y los hace intercambiables. La estrategia permite que el algoritmo varíe independientemente de los clientes que lo utilizan.
## 14. Estado (State)
El patrón Estado permite a un objeto alterar su comportamiento cuando su estado interno cambia. El objeto parecerá cambiar su clase.
## 15. Visitante (Visitor)
El patrón Visitante permite añadir operaciones a objetos sin modificar las clases de esos objetos. Proporciona una manera de separar un algoritmo de la estructura de objeto sobre la cual opera.
## 16. Mediador (Mediator)
El patrón Mediador define un objeto que encapsula cómo un conjunto de objetos interactúa. Promueve el bajo acoplamiento al evitar que los objetos se refieran entre sí explícitamente, y permite variar su interacción independientemente.
## 17. Memento
El patrón Memento captura y externaliza el estado interno de un objeto sin violar la encapsulación, de modo que el objeto puede ser restaurado a este estado más tarde. Es especialmente útil para la implementación de características de deshacer.
## 18. Composite
El patrón Composite compone objetos en estructuras de árbol para representar jerarquías parte-todo. Permite a los clientes tratar de manera uniforme objetos individuales y composiciones de objetos.
## 19. Proxy
El patrón Proxy proporciona un sustituto o marcador de posición para otro objeto para controlar el acceso a este. Puede ser utilizado para controlar operaciones costosas o para proporcionar una interfaz más simple a un conjunto complejo de clases.



