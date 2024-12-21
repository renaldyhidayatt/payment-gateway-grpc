import 'package:flutter/material.dart';

class TransferHistoryPage extends StatelessWidget {
  final List<Map<String, String>> transferHistory = [
    {
      "from": "Bank A",
      "to": "John Doe",
      "amount": "Rp 200.000",
      "date": "20 Des 2024",
    },
    {
      "from": "Wallet",
      "to": "Jane Smith",
      "amount": "Rp 150.000",
      "date": "18 Des 2024",
    },
    {
      "from": "Bank B",
      "to": "Mark Lee",
      "amount": "Rp 300.000",
      "date": "15 Des 2024",
    },
    {
      "from": "Bank A",
      "to": "Lucy Hale",
      "amount": "Rp 100.000",
      "date": "10 Des 2024",
    },
  ];

  TransferHistoryPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[50],
      appBar: AppBar(
        title: const Text(
          'Transfer History',
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.w600),
        ),
        backgroundColor: Colors.teal,
        elevation: 0,
      ),
      body: Column(
        children: [
          // Box dengan deskripsi cara transfer
          Container(
            margin: const EdgeInsets.all(16),
            padding: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(16),
              boxShadow: [
                BoxShadow(
                  color: Colors.grey.withOpacity(0.3),
                  blurRadius: 12,
                  offset: const Offset(0, 6),
                ),
              ],
            ),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text(
                  'Cara Transfer',
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                    color: Colors.teal,
                  ),
                ),
                const SizedBox(height: 8),
                const Text(
                  '1. Pilih metode pembayaran pada menu transfer.\n'
                  '2. Masukkan nama penerima dan jumlah transfer.\n'
                  '3. Tekan tombol "Kirim".\n'
                  '4. Pastikan data sudah benar, lalu konfirmasi.',
                  style: TextStyle(fontSize: 14, color: Colors.grey),
                ),
                const SizedBox(height: 12),
                Center(
                  child: ElevatedButton.icon(
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.teal,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                    ),
                    icon: const Icon(Icons.send_rounded, size: 18),
                    label: const Text('Lakukan Transfer'),
                    onPressed: () {
                      Navigator.pushNamed(context, '/transfer');
                    },
                  ),
                ),
              ],
            ),
          ),

          // Riwayat Transfer
          Expanded(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16),
              child: ListView.builder(
                itemCount: transferHistory.length,
                itemBuilder: (context, index) {
                  final transfer = transferHistory[index];
                  return _buildTransferCard(
                    from: transfer["from"]!,
                    to: transfer["to"]!,
                    amount: transfer["amount"]!,
                    date: transfer["date"]!,
                  );
                },
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildTransferCard({
    required String from,
    required String to,
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
          // Icon Transfer
          Container(
            padding: const EdgeInsets.all(12),
            decoration: BoxDecoration(
              color: Colors.teal.withOpacity(0.1),
              borderRadius: BorderRadius.circular(12),
            ),
            child: const Icon(
              Icons.swap_horiz_rounded,
              size: 28,
              color: Colors.teal,
            ),
          ),
          const SizedBox(width: 16),
          // Detail Transfer
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'From: $from',
                  style: TextStyle(
                    fontSize: 14,
                    color: Colors.grey[700],
                    fontWeight: FontWeight.w500,
                  ),
                ),
                Text(
                  'To: $to',
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
          // Amount
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
