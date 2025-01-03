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
import CardPage from '@/pages/admin/card/card';
import MerchantPage from '@/pages/admin/merchant/merchant';
import TopupPage from '@/pages/admin/topup/topup';
import TransactionPage from '@/pages/admin/transaction/transaction';
import WithdrawPage from '@/pages/admin/withdraw/withdraw';
import UserPage from '@/pages/admin/user/user';
import TransferPage from '@/pages/admin/transfer/transfer';
import RegisterPage from '@/pages/auth/register';

const router = createBrowserRouter([
  {
    path: '/',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <DashboardPage />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'card',
        element: <CardDashboard />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'merchant',
        element: <DashboardMerchant />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'saldo',
        element: <DashboardSaldo />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'topup',
        element: <DashboardTopups />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'transaction',
        element: <DashboardTransactions />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'transfers',
        element: <DashboardTransfers />,
        errorElement: <ErrorPage />,
      },
      {
        path: 'withdraws',
        element: <DashboardWithdraws />,
        errorElement: <ErrorPage />,
      },
    ],
  },
  {
    path: '/cards',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <CardPage />,
      },
    ],
  },
  {
    path: '/merchants',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <MerchantPage />,
      },
    ],
  },
  {
    path: '/saldos',
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
    path: '/topups',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <TopupPage />,
      },
    ],
  },
  {
    path: '/transactions',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <TransactionPage />,
      },
    ],
  },
  {
    path: '/transfers',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <TransferPage />,
      },
    ],
  },
  {
    path: '/users',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <UserPage />,
      },
    ],
  },
  {
    path: '/withdraws',
    element: <LayoutAdmin />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <WithdrawPage />,
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
    path: '/auth',
    element: <AuthLayout />,
    children: [
      {
        path: 'register',
        element: <RegisterPage />,
      },
      {
        path: 'login',
        element: <LoginPage />,
      },
    ],
  },
]);

export default router;
