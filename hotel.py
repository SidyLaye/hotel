import gi
gi.require_version('Gtk', '3.0')
from gi.repository import Gtk, Gdk


class Form(Gtk.Window):

    def __init__(self):
        Gtk.Window.__init__(self, title="Hotellerie")

        # Créer une grille pour les champs de texte
        grid = Gtk.Grid()
        grid.set_column_spacing(30)
        grid.set_row_spacing(30)
        self.add(grid)


        css_provider = Gtk.CssProvider()
        css = """
        #GtkLabel{
             font-size: 20px;
             font-weight: bold;
        }


        #GtkEntry {
             border: 4px solid gray;
             border-radius: 5px;
             padding: 5px;
        }

         #grid-container {
            position: center;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            margin: auto;
         }
        """


        screen = Gdk.Screen.get_default()
        style_context = Gtk.StyleContext()
        style_context.add_provider_for_screen(screen, css_provider, Gtk.STYLE_PROVIDER_PRIORITY_USER)


        grid_container = Gtk.Box(orientation=Gtk.Orientation.HORIZONTAL, spacing=10)
        grid_container.set_name("grid-container")
        grid_container.pack_start(grid, False, False, 0)
        self.add(grid_container)

        # Champ de texte pour le nom
        name = Gtk.Label(label="Nom:")
        grid.attach(name, 0, 0, 1, 1)

        name = Gtk.Entry()
        grid.attach(name, 1, 0, 1, 1)
       
        address = Gtk.Label(label="Adresse:")
        grid.attach(address, 0, 2, 1, 1)

        address = Gtk.Entry()
        grid.attach(address, 1, 2, 1, 1)

        # Champ de texte pour le numéro de téléphone
        phone = Gtk.Label(label="Téléphone:")
        grid.attach(phone, 0, 3, 1, 1)

        phone = Gtk.Entry()
        grid.attach(phone, 1, 3, 1, 1)

        # Bouton Envoyer
        send = Gtk.Button(label="Envoyer")
        send.connect("clicked", self.on_button_send_clicked, name, address, phone)
        grid.attach(send, 0, 4, 2, 1)

    def on_button_send_clicked(self, widget, name, address, phone):
        # Récupérer les valeurs des champs de texte
        name = name.get_text()
        address = address.get_text()
        phone = phone.get_text()

       
        import subprocess
        subprocess.call(["python", "accueil.py"])

        fixed = Gtk.fixed()


        image= Gtk.Image.new_from_file("C:/msys64/home/hp/imhotel.jpg")


        image.set_size_request(800,600)
        fixed.put(image, 0, 0)
        self.add(fixed)


        
        # image = Gtk.image()
        # pixbuf = GdkPixbuf.Pixbuf.new_from_file("C:/msys64/home/hp/imhotel.jpg")
        # image.set_from_pixbuf(pixbuf)


win = Form()
win.connect("destroy", Gtk.main_quit)
win.set_default_size(800,600)
win.show_all() 
Gtk.main()
