import requests
import json
import gi
gi.require_version('Gtk', '3.0')
from gi.repository import Gtk

# Écrire des données dans l'API GO
data_to_write = {'id': 1, 'name': 'John', 'value': 2.5}
response = requests.post('https://api.go.com/data', json.dumps(data_to_write))

# Lire les données de l'API GO
response = requests.get('https://api.go.com/data')
data = response.json()

# Transformer les données en une liste de dictionnaires
rows = [{'id': d['id'], 'name': d['name'], 'value': d['value']} for d in data]

# Créer le modèle de tableau
model = Gtk.ListStore(int, str, float)
model.set_column_types(int, str, float)

# Ajouter les données au modèle de tableau
for row in rows:
    model.append(row.values())

# Créer la vue de tableau
view = Gtk.TreeView(model)

# Ajouter les colonnes au tableau
for i, column_title in enumerate(['ID', 'Nom', 'Valeur']):
    renderer = Gtk.CellRendererText()
    column = Gtk.TreeViewColumn(column_title, renderer, text=i)
    view.append_column(column)

# Afficher la vue de tableau
window = Gtk.Window()
window.add(view)
window.show_all()
Gtk.main()
