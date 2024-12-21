import { CreditCard, Wallet, TrendingUp, TrendingDown } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardSaldo() {
  const barChartOptions = {
    chart: {
      id: 'balance-distribution-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Total Balance', 'Withdrawn'],
    },
    colors: ['#6366F1', '#EF4444'],
    plotOptions: {
      bar: {
        borderRadius: 5,
        columnWidth: '45%',
      },
    },
  };

  const barChartSeries = [
    {
      name: 'Amount',
      data: [5000, 1200],
    },
  ];

  const doughnutChartOptions = {
    chart: {
      id: 'transaction-status-chart',
    },
    labels: ['Completed', 'Pending', 'Failed'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const doughnutChartSeries = [70, 20, 10];

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      <div className="grid grid-cols-1 gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Balance</CardTitle>
            <Wallet className="h-6 w-6 text-gray-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">$5000</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Withdrawn</CardTitle>
            <TrendingDown className="h-6 w-6 text-gray-500" />{' '}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">$1200</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Completed Transactions
            </CardTitle>
            <TrendingUp className="h-6 w-6 text-gray-500" />{' '}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">70</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Failed Transactions
            </CardTitle>
            <CreditCard className="h-6 w-6 text-gray-500" />{' '}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">10</div>
          </CardContent>
        </Card>
      </div>

      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Balance Overview</CardTitle>
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

        <Card>
          <CardHeader>
            <CardTitle>Transaction Status</CardTitle>
          </CardHeader>
          <CardContent>
            <Chart
              options={doughnutChartOptions}
              series={doughnutChartSeries}
              type="donut"
              height={300}
            />
          </CardContent>
        </Card>
      </div>

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
                  <th className="p-2 text-left">Status</th>
                  <th className="p-2 text-left">Date</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#1001</td>
                  <td className="p-2">**** **** **** 1234</td>
                  <td className="p-2">$200</td>
                  <td className="p-2 text-green-600">Completed</td>
                  <td className="p-2">2024-12-20</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#1002</td>
                  <td className="p-2">**** **** **** 5678</td>
                  <td className="p-2">$100</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                  <td className="p-2">2024-12-21</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#1003</td>
                  <td className="p-2">**** **** **** 9101</td>
                  <td className="p-2">$300</td>
                  <td className="p-2 text-red-600">Failed</td>
                  <td className="p-2">2024-12-21</td>
                </tr>
              </tbody>
            </table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
