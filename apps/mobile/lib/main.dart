import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:mobile/components/splash.dart';
import 'package:mobile/pages/auth/registerpage.dart';
import 'package:mobile/pages/home.dart';
import 'package:mobile/pages/auth/loginpage.dart';
import 'package:mobile/pages/profilepage.dart';
import 'package:mobile/pages/topup/topup.dart';
import 'package:mobile/pages/topup/topupdetail.dart';
import 'package:mobile/pages/transfer/transfer.dart';
import 'package:mobile/pages/transfer/transferdetail.dart';
import 'package:mobile/pages/withdraw/withdraw.dart';
import 'package:mobile/pages/withdraw/withdrawdetail.dart';

void main() {
  runApp(const MyApp());
}

final GoRouter _router = GoRouter(
  initialLocation: '/',
  routes: [
    GoRoute(
      path: '/',
      builder: (context, state) => const SplashScreen(),
    ),
    GoRoute(
        name: "home",
        path: "/home",
        builder: (context, state) => const HomePage()),
    GoRoute(
      name: 'profile',
      path: '/profile',
      builder: (context, state) => const ProfilePage(),
    ),
    GoRoute(
      name: 'register',
      path: '/register',
      builder: (context, state) => const RegisterPage(),
    ),
    GoRoute(
      name: "transfer",
      path: "/transfer",
      builder: (context, state) => TransferHistoryPage(),
    ),
    GoRoute(
      name: "transferdetail",
      path: "/transferdetail",
      builder: (context, state) => TransferDetailPage(),
    ),
    GoRoute(
      name: "topup",
      path: "/topup",
      builder: (context, state) => TopupHistoryPage(),
    ),
    GoRoute(
      name: "topupdetail",
      path: "/topupdetail",
      builder: (context, state) => TopupDetailPage(),
    ),
    GoRoute(
      name: "withdraw",
      path: "/withdraw",
      builder: (context, state) => WithdrawHistoryPage(),
    ),
    GoRoute(
      name: "withdrawdetail",
      path: "/withdrawdetail",
      builder: (context, state) => WithdrawDetailPage(),
    ),
    GoRoute(
      name: 'login',
      path: '/login',
      builder: (context, state) => const LoginPage(),
    ),
  ],
);

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      routerConfig: _router,
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
    );
  }
}
