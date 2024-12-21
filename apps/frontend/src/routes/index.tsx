import DashboardPage from '@/pages/admin/dashboard/dashboard';
import LayoutAdmin from '@/layouts/admin';
import ErrorPage from '@/pages/error';
import { createBrowserRouter } from 'react-router-dom';
import AuthLayout from '@/layouts/auth';
import LoginPage from '@/pages/auth/login';
import PointOfSalePage from '@/pages/admin/point-of-sale';
import TablePayment from '@/pages/admin/table';
import TermsOfServicePage from '@/pages/TermOfService';
import CompanyProfile from '@/pages/Home';
import CardDashboard from '@/pages/admin/dashboard/dashboard-card';
import DashboardSaldo from '@/pages/admin/dashboard/dashboard-saldo';
import DashboardMerchant from '@/pages/admin/dashboard/dashboard-merchant';
import DashboardTopups from '@/pages/admin/dashboard/dashboard-topup';
import DashboardTransactions from '@/pages/admin/dashboard/dashboard-transaction';
import DashboardTransfers from '@/pages/admin/dashboard/dashboard-transfer';
import DashboardWithdraws from '@/pages/admin/dashboard/dashboard-withdraw';
import ProfilePage from '@/pages/admin/user-profile';
import SaldoPage from '@/pages/admin/saldo/saldo';

const router = createBrowserRouter([
  {
    path: '/',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <DashboardPage />,
      },
      {
        path: 'card',
        element: <CardDashboard />,
      },
      {
        path: 'merchant',
        element: <DashboardMerchant />,
      },
      {
        path: 'saldo',
        element: <DashboardSaldo />,
      },
      {
        path: 'topup',
        element: <DashboardTopups />,
      },
      {
        path: 'transaction',
        element: <DashboardTransactions />,
      },
      {
        path: 'transfers',
        element: <DashboardTransfers />,
      },
      {
        path: 'withdraws',
        element: <DashboardWithdraws />,
      },
    ],
  },
  {
    path: '/saldo',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <SaldoPage />,
      },
    ],
  },
  {
    path: '/profile',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <ProfilePage />,
      },
    ],
  },
  {
    path: '/point-of-sale',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <PointOfSalePage />,
      },
    ],
  },
  {
    path: '/term',
    element: <TermsOfServicePage />,
  },
  {
    path: '/home',
    element: <CompanyProfile />,
  },
  {
    path: '/table',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <TablePayment />,
      },
    ],
  },
  {
    path: '/login',
    element: <AuthLayout />,
    children: [
      {
        index: true,
        element: <LoginPage />,
      },
    ],
  },
]);

export default router;
