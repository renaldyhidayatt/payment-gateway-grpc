import { Repeat, FileText } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardTransactions() {
  const barChartOptions = {
    chart: {
      id: 'payment-method-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Credit Card', 'Bank Transfer', 'E-Wallet'],
    },
    colors: ['#6366F1', '#22C55E', '#F59E0B'],
    plotOptions: {
      bar: {
        borderRadius: 5,
        columnWidth: '45%',
      },
    },
  };

  const barChartSeries = [
    {
      name: 'Transaction Amount',
      data: [10000, 5000, 7000],
    },
  ];

  const pieChartOptions = {
    chart: {
      id: 'transaction-status-chart',
    },
    labels: ['Success', 'Pending', 'Failed'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const pieChartSeries = [65, 25, 10]; // Example percentages for each status

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      {/* Grid Statistik */}
      <div className="grid grid-cols-1 gap-4 md:grid-cols-4">
        {/* Total Transactions */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Transactions
            </CardTitle>
            <FileText className="h-6 w-6 text-gray-500" /> {/* Icon FileText */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">200</div>{' '}
            {/* Example total transactions */}
          </CardContent>
        </Card>

        {/* Total Success */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Transactions Success
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">130</div>{' '}
            {/* Example successful transactions */}
          </CardContent>
        </Card>

        {/* Total Pending */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Transactions Pending
            </CardTitle>
            <FileText className="h-6 w-6 text-gray-500" /> {/* Icon FileText */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">50</div>{' '}
            {/* Example pending transactions */}
          </CardContent>
        </Card>

        {/* Total Failed */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Transactions Failed
            </CardTitle>
            <FileText className="h-6 w-6 text-gray-500" /> {/* Icon FileText */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">20</div>{' '}
            {/* Example failed transactions */}
          </CardContent>
        </Card>
      </div>

      {/* Bar and Pie Charts */}
      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Transaction Amount by Payment Method</CardTitle>
          </CardHeader>
          <CardContent>
            <Chart
              options={barChartOptions}
              series={barChartSeries}
              type="bar"
              height={300}
            />
          </CardContent>
        </Card>

        {/* Pie Chart for Transaction Status */}
        <Card>
          <CardHeader>
            <CardTitle>Transaction Status</CardTitle>
          </CardHeader>
          <CardContent>
            <Chart
              options={pieChartOptions}
              series={pieChartSeries}
              type="donut"
              height={300}
            />
          </CardContent>
        </Card>
      </div>

      {/* Grid Table for Transaction Data */}
      <Card>
        <CardHeader>
          <CardTitle>Recent Transactions</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-auto">
            <table className="w-full border-collapse">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 text-left">Transaction ID</th>
                  <th className="p-2 text-left">Card Number</th>
                  <th className="p-2 text-left">Amount</th>
                  <th className="p-2 text-left">Payment Method</th>
                  <th className="p-2 text-left">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12345</td>
                  <td className="p-2">1234 5678 9876 5432</td>
                  <td className="p-2">$1000</td>
                  <td className="p-2">Credit Card</td>
                  <td className="p-2 text-green-600">Success</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12346</td>
                  <td className="p-2">1234 5678 8765 4321</td>
                  <td className="p-2">$1500</td>
                  <td className="p-2">Bank Transfer</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12347</td>
                  <td className="p-2">1234 5678 6543 2109</td>
                  <td className="p-2">$2000</td>
                  <td className="p-2">E-Wallet</td>
                  <td className="p-2 text-red-600">Failed</td>
                </tr>
              </tbody>
            </table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
