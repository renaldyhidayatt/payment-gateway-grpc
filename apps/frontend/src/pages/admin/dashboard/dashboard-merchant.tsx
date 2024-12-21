import { Store } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Chart from 'react-apexcharts';

export default function DashboardMerchant() {
  const barChartOptions = {
    chart: {
      id: 'merchant-status-bar-chart',
      toolbar: { show: false },
    },
    xaxis: {
      categories: ['Active', 'Inactive', 'Suspended'],
    },
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    plotOptions: {
      bar: {
        borderRadius: 5,
        columnWidth: '45%',
      },
    },
  };

  const barChartSeries = [
    {
      name: 'Merchants',
      data: [30, 10, 5],
    },
  ];

  const pieChartOptions = {
    chart: {
      id: 'merchant-status-pie-chart',
    },
    labels: ['Active', 'Inactive', 'Suspended'],
    colors: ['#22C55E', '#EAB308', '#EF4444'],
    legend: { position: 'bottom' },
  };

  const pieChartSeries = [30, 10, 5];

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      <div className="grid grid-cols-1 gap-4 md:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Merchants
            </CardTitle>
            <Store className="h-6 w-6 text-gray-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">45</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Active Merchants
            </CardTitle>
            <Store className="h-6 w-6 text-green-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">30</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Suspended Merchants
            </CardTitle>
            <Store className="h-6 w-6 text-red-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">5</div>
          </CardContent>
        </Card>
      </div>

      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Merchant Status (Bar Chart)</CardTitle>
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
            <CardTitle>Merchant Status (Pie Chart)</CardTitle>
          </CardHeader>
          <CardContent>
            <Chart
              options={pieChartOptions}
              series={pieChartSeries}
              type="pie"
              height={300}
            />
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Recent Merchants</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="overflow-auto">
            <table className="w-full border-collapse">
              <thead>
                <tr className="bg-gray-100">
                  <th className="p-2 text-left">Merchant ID</th>
                  <th className="p-2 text-left">Name</th>
                  <th className="p-2 text-left">API Key</th>
                  <th className="p-2 text-left">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#1</td>
                  <td className="p-2">Merchant A</td>
                  <td className="p-2">API_KEY_123</td>
                  <td className="p-2 text-green-600">Active</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#2</td>
                  <td className="p-2">Merchant B</td>
                  <td className="p-2">API_KEY_456</td>
                  <td className="p-2 text-red-600">Suspended</td>
                </tr>
                <tr className="even:bg-gray-50">
                  <td className="p-2">#3</td>
                  <td className="p-2">Merchant C</td>
                  <td className="p-2">API_KEY_789</td>
                  <td className="p-2 text-yellow-600">Inactive</td>
                </tr>
              </tbody>
            </table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
