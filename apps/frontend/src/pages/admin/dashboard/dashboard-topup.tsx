import { Users, DollarSign, Repeat, FileText } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardTopups() {
  const barChartOptions = {
    chart: {
      id: 'topup-method-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Credit Card', 'Bank Transfer', 'E-Wallet'], // Top-up methods
    },
    colors: ['#6366F1', '#22C55E', '#F59E0B'], // Color palette for the bars
    plotOptions: {
      bar: {
        borderRadius: 5,
        columnWidth: '45%',
      },
    },
  };

  const barChartSeries = [
    {
      name: 'Topup Amount',
      data: [5000, 3000, 2000], // Sample top-up amounts for each method
    },
  ];

  const pieChartOptions = {
    chart: {
      id: 'topup-status-chart',
    },
    labels: ['Success', 'Pending', 'Failed'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const pieChartSeries = [70, 20, 10]; // Sample top-up statuses (percentage)

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      {/* Grid Statistik */}
      <div className="grid grid-cols-1 gap-4 md:grid-cols-4">
        {/* Total Topups */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Topups</CardTitle>
            <DollarSign className="h-6 w-6 text-gray-500" />{' '}
            {/* Icon Dollar Sign */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">150</div>{' '}
            {/* Sample total top-ups */}
          </CardContent>
        </Card>

        {/* Total Success */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Topups Success
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">105</div>{' '}
            {/* Sample successful top-ups */}
          </CardContent>
        </Card>

        {/* Total Pending */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Topups Pending
            </CardTitle>
            <FileText className="h-6 w-6 text-gray-500" />{' '}
            {/* Icon File Text */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">30</div>{' '}
            {/* Sample pending top-ups */}
          </CardContent>
        </Card>

        {/* Total Failed */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Topups Failed</CardTitle>
            <FileText className="h-6 w-6 text-gray-500" />{' '}
            {/* Icon File Text */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">15</div>{' '}
            {/* Sample failed top-ups */}
          </CardContent>
        </Card>
      </div>

      {/* Bar and Pie Charts */}
      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Topup Method Overview</CardTitle>
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

        {/* Pie Chart for Topup Status */}
        <Card>
          <CardHeader>
            <CardTitle>Topup Status</CardTitle>
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

      {/* Grid Table for Topup Transactions */}
      <Card>
        <CardHeader>
          <CardTitle>Recent Topup Transactions</CardTitle>
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
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12345</td>
                  <td className="p-2">1234 5678 9876 5432</td>
                  <td className="p-2">$500</td>
                  <td className="p-2 text-green-600">Success</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12346</td>
                  <td className="p-2">1234 5678 8765 4321</td>
                  <td className="p-2">$1000</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12347</td>
                  <td className="p-2">1234 5678 6543 2109</td>
                  <td className="p-2">$250</td>
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
