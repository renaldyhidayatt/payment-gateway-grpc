import { CreditCard, ShieldCheck } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardCard() {
  const barChartOptions = {
    chart: {
      id: 'card-type-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Visa', 'MasterCard', 'Amex'],
    },
    colors: ['#6366F1', '#22C55E', '#EAB308'],
    plotOptions: {
      bar: {
        borderRadius: 5,
        columnWidth: '45%',
      },
    },
  };

  const barChartSeries = [
    {
      name: 'Cards',
      data: [80, 50, 20],
    },
  ];

  // Donut Chart Options & Data
  const doughnutChartOptions = {
    chart: {
      id: 'card-status-chart',
    },
    labels: ['Active', 'Pending', 'Expired'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const doughnutChartSeries = [120, 20, 10];

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Cards</CardTitle>
            <CreditCard className="h-6 w-6 text-gray-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">150</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Active Cards</CardTitle>
            <ShieldCheck className="h-6 w-6 text-gray-500" />{' '}
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">120</div>
          </CardContent>
        </Card>
      </div>

      {/* Charts */}
      <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Card Types Overview</CardTitle>
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
            <CardTitle>Card Status Distribution</CardTitle>
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
          <CardTitle>Recent Cards</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-auto">
            <table className="w-full border-collapse">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 text-left">Card ID</th>
                  <th className="p-2 text-left">Card Number</th>
                  <th className="p-2 text-left">Card Type</th>
                  <th className="p-2 text-left">Expire Date</th>
                  <th className="p-2 text-left">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#101</td>
                  <td className="p-2">**** **** **** 1234</td>
                  <td className="p-2">Visa</td>
                  <td className="p-2">12/2025</td>
                  <td className="p-2 text-green-600">Active</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#102</td>
                  <td className="p-2">**** **** **** 5678</td>
                  <td className="p-2">MasterCard</td>
                  <td className="p-2">05/2024</td>
                  <td className="p-2 text-yellow-600">Pending</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#103</td>
                  <td className="p-2">**** **** **** 9101</td>
                  <td className="p-2">Amex</td>
                  <td className="p-2">09/2026</td>
                  <td className="p-2 text-red-600">Expired</td>
                </tr>
              </tbody>
            </table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
