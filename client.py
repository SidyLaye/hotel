import requests
import json
import gi
gi.require_version('Gtk', '3.0')
from gi.repository import Gtk

#   Écrire des données dans l'API GO
# data_to_write = {'id': 1, 'name': 'John', 'value': 2.5}
# response = requests.post('localhost:11111/chambres', json.dumps(data_to_write))

# Lire les données de l'API GO
response = requests.get('localhost:11111/clients')
data = response.json()

# Transformer les données en une liste de dictionnaires
rows = [{'nom': d['nom'], 'prenom': d['prenom'], 'telephone': d['telephone'],  'date_reservation':d['date_reservation'], 'date_entree':d['date_entree'],
          'date_sortie':d['date_sortie'],'nuite' :d['nuite'], 'tarif_chambre' : d['tarif_chambre'], 'tarif_serAnnexe': d['tarif_serAnnexe'], 'total' : d['total'] } for d in data]

# Créer le modèle de tableau
model = Gtk.ListStore(int, int, str)
model.set_column_types(int, int, str)

# Ajouter les données au modèle de tableau
for row in rows:
    model.append(row.values())

# Créer la vue de tableau
view = Gtk.TreeView(model)

# Ajouter les colonnes au tableau
for i, column_title in enumerate(['nom', 'prenom', 'telephone', 'date_reservation', 'date_entree', 'date_sortie','nuite', 'tarif_chambre', 'tarif_serAnnexe', 'total']):
    renderer = Gtk.CellRendererText()
    column = Gtk.TreeViewColumn(column_title, renderer, text=i)
    view.append_column(column)

# Afficher la vue de tableau
window = Gtk.Window()
window.add(view)
window.show_all()
Gtk.main()
