import { Header } from '@/components/admin/header';
import { AppSidebar } from '@/components/app-sidebar';
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar';
import { Outlet } from 'react-router-dom';

function LayoutAdmin() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <Header />
        <br />
        <Outlet />
      </SidebarInset>
    </SidebarProvider>
  );
}

export default LayoutAdmin;
