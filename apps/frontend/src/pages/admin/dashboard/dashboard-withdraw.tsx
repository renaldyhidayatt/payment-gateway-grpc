import { Users, DollarSign, Repeat, FileText } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardWithdraws() {
  const barChartOptions = {
    chart: {
      id: 'withdraw-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Withdraws', 'Amount'],
    },
    colors: ['#6366F1', '#22C55E'],
    plotOptions: {
      bar: {
        borderRadius: 5,
        columnWidth: '45%',
      },
    },
  };

  const barChartSeries = [
    {
      name: 'Count',
      data: [150, 120], // Example data: 150 withdrawals, $120 amount
    },
  ];

  const doughnutChartOptions = {
    chart: {
      id: 'withdraw-status-chart',
    },
    labels: ['Success', 'Pending', 'Failed'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const doughnutChartSeries = [60, 25, 15]; // Example: 60% success, 25% pending, 15% failed

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      {/* Grid Statistik */}
      <div className="grid grid-cols-1 gap-4 md:grid-cols-4">
        {/* Total Withdraws */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Withdraws
            </CardTitle>
            <DollarSign className="h-6 w-6 text-gray-500" />{' '}
            {/* Icon Withdraws */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">150</div>
          </CardContent>
        </Card>

        {/* Total Amount Withdrawn */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Amount Withdrawn
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">$1200</div>
          </CardContent>
        </Card>

        {/* Pending Withdrawals */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Pending Withdrawals
            </CardTitle>
            <FileText className="h-6 w-6 text-gray-500" /> {/* Icon FileText */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">25</div>
          </CardContent>
        </Card>
      </div>

      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Withdraw Overview</CardTitle>
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

        {/* Doughnut Chart for Withdraw Status */}
        <Card>
          <CardHeader>
            <CardTitle>Withdraw Status</CardTitle>
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

      {/* Grid Table for Recent Withdrawals */}
      <Card>
        <CardHeader>
          <CardTitle>Recent Withdrawals</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-auto">
            <table className="w-full border-collapse">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 text-left">Withdrawal ID</th>
                  <th className="p-2 text-left">User</th>
                  <th className="p-2 text-left">Amount</th>
                  <th className="p-2 text-left">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#98765</td>
                  <td className="p-2">John Doe</td>
                  <td className="p-2">$500</td>
                  <td className="p-2 text-green-600">Success</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#98766</td>
                  <td className="p-2">Jane Smith</td>
                  <td className="p-2">$200</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#98767</td>
                  <td className="p-2">Alice Brown</td>
                  <td className="p-2">$300</td>
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
