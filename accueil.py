import gi
gi.require_version('Gtk', '3.0')
from gi.repository import Gtk, Gdk
import subprocess

class MenuExample(Gtk.Window):

    def __init__(self):
        Gtk.Window.__init__(self, title="ACCUEIL HOTEL")
        
        
        vbox = Gtk.Box(orientation=Gtk.Orientation.VERTICAL, spacing=0)
        self.add(vbox)


        titre = Gtk.Label()
        titre.set_text("Bienvenue dans votre espace hôtel !")
        vbox.pack_start(titre, False, False, 0)


        # Création d'une barre de menus
        menubar = Gtk.MenuBar()
        menubar.set_hexpand(True)
        menubar.set_vexpand(False)
        vbox.pack_start(menubar, False, False, 0)
        self.add(menubar)


        css_provider = Gtk.CssProvider()
        css = """
            
        """
        screen = Gdk.Screen.get_default()
        style_context = Gtk.StyleContext()
        style_context.add_provider_for_screen(screen, css_provider, Gtk.STYLE_PROVIDER_PRIORITY_USER)


        # Création d'un élément de menu "Information" avec des sous-options
        infomenu = Gtk.Menu()
        info_item = Gtk.MenuItem(label="Informations")
        info_item.set_submenu(infomenu)

        chambre = Gtk.Menu()
        chambre_item = Gtk.MenuItem(label="Chambre")
        chambre_item.set_submenu(chambre)
        infomenu.append(chambre_item)

        # Sous éléments de "Chambre"
        consulter_item1 = Gtk.MenuItem(label="Consulter chambres libres")
        consulter_item1.connect("activate", self.on_menu_chambre)
        chambre.append(consulter_item1)

        consulter_item2 = Gtk.MenuItem(label="Consulter chambres occupées")
        consulter_item2.connect("activate", self.on_menu_chambre)
        chambre.append(consulter_item2)

        consulter_item3 = Gtk.MenuItem(label="Consulter chambres réservées")
        consulter_item3.connect("activate", self.on_menu_chambre)
        chambre.append(consulter_item3)


        #Sous éléments de "Clients"
        client= Gtk.Menu()
        client_item = Gtk.MenuItem(label="Clients")
        client_item.set_submenu(client)
        infomenu.append(client_item)

       
        client.append(Gtk.SeparatorMenuItem())

        cp_item = Gtk.MenuItem(label="Clients présents")
        cp_item.connect("activate", self.on_menu_client)
        client.append(cp_item)

        cr_item = Gtk.MenuItem(label="Clients réservés")
        cr_item.connect("activate", self.on_menu_client)
        client.append(cr_item)

        cs_item = Gtk.MenuItem(label="Clients sortants")
        cs_item.connect("activate", self.on_menu_client)
        client.append(cs_item)


        #Sous éléments de "Réservation"
        reservation = Gtk.Menu()
        reservation_item = Gtk.MenuItem(label="Réservations")
        reservation_item.set_submenu(reservation)
        infomenu.append(reservation_item)

        consres_item = Gtk.MenuItem(label="Consulter résevation")
        consres_item.connect("activate", self.on_menu_reservation)
        reservation.append(consres_item)


        # Sous éléments de "Factures"
        facture = Gtk.Menu()
        facture_item = Gtk.MenuItem(label="Factures")
        facture_item.set_submenu(facture)
        infomenu.append(facture_item)

        fact_item = Gtk.MenuItem(label="Consulter la liste des factures")
        fact_item.connect("activate", self.on_menu_facture)
        facture.append(fact_item)

        fact_item1 = Gtk.MenuItem(label="Consulter la liste des factures d'aujourd'hui")
        fact_item.connect("activate", self.on_menu_facture)
        facture.append(fact_item1)

        gen_item = Gtk.MenuItem(label="Générer facture")
        gen_item.connect("activate", self.on_menu_facture)
        facture.append(gen_item)


        # Sous éléments de "Statistiques"
        statistique = Gtk.Menu()
        statistique_item = Gtk.MenuItem(label="Statistiques")
        statistique_item.set_submenu(statistique)
        infomenu.append(statistique_item)

        sm = Gtk.Menu()
        sm_item = Gtk.MenuItem(label="Statistique mensuelle")
        sm_item.set_submenu(sm)
        statistique.append(sm_item)

        fa_item = Gtk.MenuItem(label="Flux d'argent")
        fa_item.connect("activate", self.on_menu_statistique)
        sm.append(fa_item)

        fp_item = Gtk.MenuItem(label="Flux de personnes")
        fp_item.connect("activate", self.on_menu_statistique)
        sm.append(fp_item)

        ssa = Gtk.Menu()
        ssa_item = Gtk.MenuItem(label="Statistique semestrielle et annuelle")
        ssa_item.set_submenu(ssa)
        statistique.append(ssa_item)

        fa1_item = Gtk.MenuItem(label="Flux d'argent")
        fa1_item.connect("activate", self.on_menu_statistique)
        ssa.append(fa1_item)

        fp1_item = Gtk.MenuItem(label="Flux de personnes")
        fp1_item.connect("activate", self.on_menu_statistique)
        ssa.append(fp1_item)


        # élément "Aide"
        aide_item = Gtk.MenuItem(label="Aide")
        # aide_item.connect("activate", self.on_menu_item_clicked)
        infomenu.append(aide_item)
        
        #Elément "Quitter"
        quitter_item = Gtk.MenuItem(label="Quitter")
        quitter_item.connect("activate", Gtk.main_quit)
        infomenu.append(quitter_item)


        # Création d'un élément de menu "Modifier" avec des sous-options
        modifmenu = Gtk.Menu()
        modif_item = Gtk.MenuItem(label="Modifier")
        modif_item.set_submenu(modifmenu)

        #Sous éléments de "Chambre" dans "Modifier"
        chambre = Gtk.Menu()
        chambre_item = Gtk.MenuItem(label="Chambres")
        chambre_item.set_submenu(chambre)
        modifmenu.append(chambre_item)

        consulter1_item = Gtk.MenuItem(label="Consulter chambres libres")
        consulter1_item.connect("activate", self.on_menu_chambre)
        chambre.append(consulter1_item)

        consulter2_item = Gtk.MenuItem(label="Consulter chambres occupées")
        consulter2_item.connect("activate", self.on_menu_chambre)
        chambre.append(consulter2_item)

        consulter3_item = Gtk.MenuItem(label="Consulter chambres réservées")
        consulter3_item.connect("activate", self.on_menu_chambre)
        chambre.append(consulter3_item)

        chambre.append(Gtk.SeparatorMenuItem())

        modif1_item = Gtk.MenuItem(label="Modifier la classe de la chambre")
        modif1_item.connect("activate", self.on_menu_chambre)
        chambre.append(modif1_item)

        modif2_item = Gtk.MenuItem(label="Modifier l'état de la chambre")
        modif2_item.connect("activate", self.on_menu_chambre)
        chambre.append(modif2_item)

        
        #Sous éléments de "Clients" dans "Modifier"
        client = Gtk.Menu()
        clients_item = Gtk.MenuItem(label="Clients")
        clients_item.set_submenu(client)
        modifmenu.append(clients_item)

        enregistrer_item = Gtk.MenuItem(label="Enregistrer client")
        enregistrer_item.connect("activate", self.on_menu_client)
        client.append(enregistrer_item)

        client.append(Gtk.SeparatorMenuItem())

        cp1_item = Gtk.MenuItem(label="Clients présents")
        cp1_item.connect("activate", self.on_menu_client)
        client.append(cp1_item)

        cr1_item = Gtk.MenuItem(label="Clients réservés")
        cr1_item.connect("activate", self.on_menu_client)
        client.append(cr1_item)

        cs1_item = Gtk.MenuItem(label="Clients sortants")
        cs1_item.connect("activate", self.on_menu_client)
        client.append(cs1_item)

        client.append(Gtk.SeparatorMenuItem())

        sup_item = Gtk.MenuItem(label="Supprimer client")
        sup_item.connect("activate", self.on_menu_client)
        client.append(sup_item)


        #Sous éléments de "Réservation" dans "Modifier"
        reservation = Gtk.Menu()
        reservations_item = Gtk.MenuItem(label="Réservations")
        reservations_item.set_submenu(reservation)
        modifmenu.append(reservations_item)

        consres1_item = Gtk.MenuItem(label="Consulter résevation")
        consres1_item.connect("activate", self.on_menu_reservation)
        reservation.append(consres1_item)

        ajoutres_item = Gtk.MenuItem(label="Ajouter résevation")
        ajoutres_item.connect("activate", self.on_menu_reservation)
        reservation.append(ajoutres_item)

        sup_item = Gtk.MenuItem(label="Supprimer résevation")
        sup_item.connect("activate", self.on_menu_reservation)
        reservation.append(sup_item)

        
        #Sous éléments de "Factures" dans "Modifier"
        facture = Gtk.Menu()
        factures_item = Gtk.MenuItem(label="Factures")
        factures_item.set_submenu(facture)
        modifmenu.append(factures_item)

        fact1_item = Gtk.MenuItem(label="Consulter la liste des factures")
        fact1_item.connect("activate", self.on_menu_facture)
        facture.append(fact1_item)

        fact2_item = Gtk.MenuItem(label="Consulter la liste des factures d'aujourd'hui")
        fact2_item.connect("activate", self.on_menu_facture)
        facture.append(fact2_item)

        gen1_item = Gtk.MenuItem(label="Générer facture")
        gen1_item.connect("activate", self.on_menu_facture)
        facture.append(gen1_item)


        #Sous éléments de "Statistiques" dans "Modifier"
        statistique = Gtk.Menu()
        statistiques_item = Gtk.MenuItem(label="Statistiques")
        statistiques_item.set_submenu(statistique)
        modifmenu.append(statistiques_item)


        sm1 = Gtk.Menu()
        sm1_item = Gtk.MenuItem(label="Statistique mensuelle")
        sm1_item.set_submenu(sm1)
        statistique.append(sm1_item)

        fa_item1 = Gtk.MenuItem(label="Flux d'argent")
        # fa_item1.connect()
        sm1.append(fa_item1)

        fp_item1 = Gtk.MenuItem(label="Flux de personnes")
        # fp_item1.connect()
        sm1.append(fp_item1)

        ssa1 = Gtk.Menu()
        ssa1_item = Gtk.MenuItem(label="Statistique semestrielle et annuelle")
        ssa1_item.set_submenu(ssa1)
        statistique.append(ssa1_item)

        fa1_item1 = Gtk.MenuItem(label="Flux d'argent")
        fa1_item1.connect("activate", self.on_menu_statistique)
        ssa1.append(fa1_item1)

        fp1_item1 = Gtk.MenuItem(label="Flux de personnes")
        fp1_item1.connect("activate", self.on_menu_statistique)
        ssa1.append(fp1_item1)

        
         # Création d'un élément de menu "Reinitialiser" avec des sous-options
        reinitmenu = Gtk.Menu()
        reinit_item = Gtk.MenuItem(label="Reinitialiser")
        reinit_item.set_submenu(reinitmenu)

        reset_item = Gtk.MenuItem(label="Reset")
        reset_item.connect("activate", self.on_button_clicked)
        reinitmenu.append(reset_item)

        
         # Création d'un élément de menu "Tarif" avec des sous-options
        tarifmenu = Gtk.Menu()
        tarif_item = Gtk.MenuItem(label="Tarif")
        tarif_item.set_submenu(tarifmenu)

        tcat_item = Gtk.MenuItem(label="Tarif categorie")
        tcat_item.connect("activate", self.on_menu_categorie)
        tarifmenu.append(tcat_item)

        ps_item = Gtk.MenuItem(label="Prix service")
        ps_item.connect("activate", self.on_menu_categorie)
        tarifmenu.append(ps_item)

        
         # Création d'un élément de menu "Aide" avec des sous-options
        aidemenu = Gtk.Menu()
        aide_item = Gtk.MenuItem(label="Aide")
        aide_item.set_submenu(aidemenu)

        service_item = Gtk.MenuItem(label="Service client")
        # service_item.connect()
        aidemenu.append(service_item)


        # Ajouter les éléments  à la barre de menu
        menubar.append(info_item)
        menubar.append(modif_item)
        menubar.append(reinit_item)
        menubar.append(tarif_item)
        menubar.append(aide_item)
        

    def on_menu_chambre(self, widget):
        # print(widget.get_label() + " a été cliqué")
        # dialog = Gtk.MessageDialog(self, 0, Gtk.MessageType.INFO,
        #        buttons = Gtk.ButtonsType.OK,  text="Ceci est un exemple de popup."  )
        # dialog.run()
        # dialog.destroy()
        subprocess.run(["python3", "chambre.py"])


    def on_menu_client(self, widget):
         subprocess.run(["python3", "client.py"])


    def on_menu_facture(self, widget):
         subprocess.run(["python3", "facture.py"])


    def on_menu_reservation(self, widget):
         subprocess.run(["python3", "reservation.py"])


    def on_menu_statistique(self, widget):
         subprocess.run(["python3", "statistique.py"])


    def on_menu_categorie(self, widget):
         subprocess.run(["python3", "categorie.py"])

    
    def on_button_clicked(self, widget):
         self.destroy()


win = MenuExample()
win.connect("destroy", Gtk.main_quit)
win.show_all()
Gtk.main()
