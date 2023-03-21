import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:shop_app/screens/home/AjoutReservation.dart';
import 'package:shop_app/screens/home/ReservationFormBody.dart';
import 'package:shop_app/screens/home/components/render_button.dart';

import '../../../size_config.dart';

class Categories extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    List<Map<String, dynamic>> categories = [
      {
        "icon": "assets/icons/Flash Icon.svg",
        "text": "Chambre",
        "link": [
          new Link(
            title: "Ajout de chambres",
          ),
          new Link(title: "Liste des chambres"),
        ]
      },
      {
        "icon": "assets/icons/Bill Icon.svg",
        "text": "Réservation",
        "link": [
          new Link(
              title: "Ajout de réservation",
              direction: MaterialPageRoute(
                  builder: (context) => AjouterReservation())),
          new Link(title: "Annulation de réservation"),
          new Link(title: "Liste des réservations"),
        ]
      },
      {
        "icon": "assets/icons/Game Icon.svg",
        "text": "Facture",
        "link": [
          new Link(title: "Liste des factures"),
          new Link(title: "Annulation de factures"),
        ]
      },
      {
        "icon": "assets/icons/Gift Icon.svg",
        "text": "Statistiques",
        "link": [
          new Link(title: "Statistiques mensuelles"),
          new Link(title: "Statistiques semestrielles/annuelles"),
        ]
      },
    ];
    return Padding(
      padding: EdgeInsets.all(getProportionateScreenWidth(20)),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: List.generate(
          categories.length,
          (index) => CategoryCard(
            icon: categories[index]["icon"],
            text: categories[index]["text"],
            press: () {
              Navigator.push(
                  context,
                  MaterialPageRoute(
                      builder: (context) => RenderButton(
                          link: categories[index]["link"],
                          title: categories[index]["text"])));
            },
          ),
        ),
      ),
    );
  }
}

class CategoryCard extends StatelessWidget {
  const CategoryCard({
    Key? key,
    required this.icon,
    required this.text,
    required this.press,
  }) : super(key: key);

  final String? icon, text;
  final GestureTapCallback press;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: press,
      child: SizedBox(
        width: getProportionateScreenWidth(75),
        child: Column(
          children: [
            Container(
              padding: EdgeInsets.all(getProportionateScreenWidth(15)),
              height: getProportionateScreenWidth(55),
              width: getProportionateScreenWidth(55),
              decoration: BoxDecoration(
                color: Color(0xFFFFECDF),
                borderRadius: BorderRadius.circular(10),
              ),
              child: SvgPicture.asset(icon!),
            ),
            SizedBox(height: 4),
            Text(text!, textAlign: TextAlign.center)
          ],
        ),
      ),
    );
  }
}
