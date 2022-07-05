If we want to run the project in a way other than connecting with the http protocol (for example, using geen or gorillamux), we know that each of these is actually an application of the http protocol, but with more capabilities.
The http package helps us to use the http protocol.

In this part of the code:

type home struct {
template *template .Template
}

The first template from the structure field is not exported.
Public and private variables in GO are generally displayed in this way.
If they are written in capital letters, they are exported and can be used inside the structure.
But if they are written with lowercase letters, they are internal.
The second template is a package called template that we have exported and are using.
The third template is an arbitrary field with any arbitrary type that is exported and that is why it is called.

In this part of the code :

package viewmodel

We have created a package called viewmodel elsewhere and we want to use its features in the main file, so we need to import it.