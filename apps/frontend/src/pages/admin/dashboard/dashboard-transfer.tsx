import { Users, Store, Repeat, FileText } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardTransfers() {
  // Bar Chart Options - Transfer Amount by Payment Method
  const barChartOptions = {
    chart: {
      id: 'payment-method-transfer-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Credit Card', 'Bank Transfer', 'E-Wallet'], // Payment methods for transfers
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
      name: 'Transfer Amount',
      data: [5000, 3000, 2000], // Example transfer amounts for each method
    },
  ];

  // Doughnut Chart Options - Transfer Status
  const doughnutChartOptions = {
    chart: {
      id: 'transfer-status-chart',
    },
    labels: ['Success', 'Pending', 'Failed'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const doughnutChartSeries = [60, 30, 10]; // Example percentages for each transfer status

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      {/* Grid Statistik */}
      <div className="grid grid-cols-1 gap-4 md:grid-cols-4">
        {/* Total Transfers */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Transfers
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">80</div>{' '}
            {/* Example total transfers */}
          </CardContent>
        </Card>

        {/* Total Successful Transfers */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Successful Transfers
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">48</div>{' '}
            {/* Example successful transfers */}
          </CardContent>
        </Card>

        {/* Total Pending Transfers */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Pending Transfers
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">24</div>{' '}
            {/* Example pending transfers */}
          </CardContent>
        </Card>

        {/* Total Failed Transfers */}
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Failed Transfers
            </CardTitle>
            <Repeat className="h-6 w-6 text-gray-500" /> {/* Icon Repeat */}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">8</div>{' '}
            {/* Example failed transfers */}
          </CardContent>
        </Card>
      </div>

      {/* Bar and Pie Charts */}
      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Transfer Amount by Payment Method</CardTitle>
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

        {/* Doughnut Chart for Transfer Status */}
        <Card>
          <CardHeader>
            <CardTitle>Transfer Status</CardTitle>
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

      {/* Recent Transfers Table */}
      <Card>
        <CardHeader>
          <CardTitle>Recent Transfers</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-auto">
            <table className="w-full border-collapse">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 text-left">Transfer ID</th>
                  <th className="p-2 text-left">From Card</th>
                  <th className="p-2 text-left">To Card</th>
                  <th className="p-2 text-left">Amount</th>
                  <th className="p-2 text-left">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12345</td>
                  <td className="p-2">1234 5678 9876 5432</td>
                  <td className="p-2">1234 5678 8765 4321</td>
                  <td className="p-2">$1000</td>
                  <td className="p-2 text-green-600">Success</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12346</td>
                  <td className="p-2">1234 5678 8765 4321</td>
                  <td className="p-2">1234 5678 6543 2109</td>
                  <td className="p-2">$500</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#T12347</td>
                  <td className="p-2">1234 5678 6543 2109</td>
                  <td className="p-2">1234 5678 8765 4321</td>
                  <td className="p-2">$2000</td>
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
