import { Users, Store, Repeat, FileText } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardPage() {
  const barChartOptions = {
    chart: {
      id: 'user-transfer-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Users', 'Transfers'],
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
      data: [120, 80],
    },
  ];

  const doughnutChartOptions = {
    chart: {
      id: 'transaction-chart',
    },
    labels: ['Success', 'Pending', 'Failed'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const doughnutChartSeries = [70, 20, 10];

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      <div className="grid grid-cols-1 gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Users</CardTitle>
            <Users className="h-6 w-6 text-gray-500" /> {/* Icon Users */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">120</div>
          </CardContent>
        </Card>

        {/* Total Merchants */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Merchants
            </CardTitle>
            <Store className="h-6 w-6 text-gray-500" /> {/* Icon Store */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">45</div>
          </CardContent>
        </Card>

        {/* Total Transfers */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Transfers
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">80</div>
          </CardContent>
        </Card>

        {/* Total Transactions */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Transactions
            </CardTitle>
            <FileText className="h-6 w-6 text-gray-500" /> {/* Icon FileText */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">200</div>
          </CardContent>
        </Card>
      </div>

      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>User & Transfer Overview</CardTitle>
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

        {/* Grafik Doughnut Chart */}
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

      {/* Grid Table */}
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
                  <th className="p-2 text-left">User</th>
                  <th className="p-2 text-left">Amount</th>
                  <th className="p-2 text-left">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#12345</td>
                  <td className="p-2">John Doe</td>
                  <td className="p-2">$100</td>
                  <td className="p-2 text-green-600">Success</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#12346</td>
                  <td className="p-2">Jane Smith</td>
                  <td className="p-2">$200</td>
                  <td className="p-2 text-red-600">Failed</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#12347</td>
                  <td className="p-2">Alice Brown</td>
                  <td className="p-2">$150</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                </tr>
              </tbody>
            </table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
