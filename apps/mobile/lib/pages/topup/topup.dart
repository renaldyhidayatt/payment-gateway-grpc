import 'package:flutter/material.dart';

class TopupHistoryPage extends StatelessWidget {
  final List<Map<String, String>> topupHistory = [
    {
      "method": "Bank Transfer",
      "amount": "Rp 500.000",
      "date": "20 Des 2024",
    },
    {
      "method": "E-Wallet",
      "amount": "Rp 300.000",
      "date": "18 Des 2024",
    },
    {
      "method": "Credit Card",
      "amount": "Rp 200.000",
      "date": "15 Des 2024",
    },
    {
      "method": "Cryptocurrency",
      "amount": "Rp 1.000.000",
      "date": "10 Des 2024",
    },
  ];

  TopupHistoryPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[50],
      appBar: AppBar(
        title: const Text(
          'Topup History',
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.w600),
        ),
        backgroundColor: Colors.teal,
        elevation: 0,
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: ListView.builder(
          itemCount: topupHistory.length,
          itemBuilder: (context, index) {
            final topup = topupHistory[index];
            return _buildTopupCard(
              method: topup["method"]!,
              amount: topup["amount"]!,
              date: topup["date"]!,
            );
          },
        ),
      ),
    );
  }

  Widget _buildTopupCard({
    required String method,
    required String amount,
    required String date,
  }) {
    return Container(
      margin: const EdgeInsets.only(bottom: 16),
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(16),
        boxShadow: [
          BoxShadow(
            color: Colors.grey.withOpacity(0.2),
            blurRadius: 10,
            offset: const Offset(0, 4),
          ),
        ],
      ),
      child: Row(
        children: [
          // Icon Metode Topup
          Container(
            padding: const EdgeInsets.all(12),
            decoration: BoxDecoration(
              color: Colors.teal.withOpacity(0.1),
              borderRadius: BorderRadius.circular(12),
            ),
            child: const Icon(
              Icons.account_balance_wallet_rounded,
              size: 28,
              color: Colors.teal,
            ),
          ),
          const SizedBox(width: 16),
          // Detail Topup
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  method,
                  style: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                const SizedBox(height: 8),
                Text(
                  date,
                  style: TextStyle(
                    fontSize: 12,
                    color: Colors.grey[500],
                  ),
                ),
              ],
            ),
          ),
          // Jumlah Topup
          Text(
            amount,
            style: const TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.bold,
              color: Colors.teal,
            ),
          ),
        ],
      ),
    );
  }
}
