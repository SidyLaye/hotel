import 'package:flutter/material.dart';
import 'package:shop_app/screens/home/ReservationFormBody.dart';

class AjouterReservation extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Ajouter reservation"),
      ),
      body: ReservationFormBody(),
    );
  }
}
