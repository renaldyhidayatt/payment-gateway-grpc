import DashboardPage from '@/pages/admin/dashboard';
import LayoutAdmin from '@/layouts/admin';
import ErrorPage from '@/pages/error';
import { createBrowserRouter } from 'react-router-dom';
import AuthLayout from '@/layouts/auth';
import LoginPage from '@/pages/auth/login';
import PointOfSalePage from '@/pages/admin/point-of-sale';
import TablePayment from '@/pages/admin/table';
import TermsOfServicePage from '@/pages/TermOfService';
import CompanyProfile from '@/pages/Home';

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
