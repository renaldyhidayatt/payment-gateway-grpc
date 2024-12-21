import { Header } from '@/components/admin/header';
import { AppSidebar } from '@/components/app-sidebar';
import SplashScreen from '@/components/splash';
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar';
import { useEffect, useState } from 'react';
import { Outlet, useNavigation } from 'react-router-dom';

function LayoutAdmin() {
  const [isLoading, setIsLoading] = useState(true);
  const navigation = useNavigation();

  useEffect(() => {
    console.log('Navigation state:', navigation.state);
    if (navigation.state === 'loading') {
      setIsLoading(true);
    } else {
      const timer = setTimeout(() => setIsLoading(false), 1500);

      return () => clearTimeout(timer);
    }
  }, [navigation.state]);

  return (
    <>
      {isLoading ? (
        <SplashScreen />
      ) : (
        <SidebarProvider>
          <AppSidebar />
          <SidebarInset>
            <Header />
            <br />
            <Outlet />
          </SidebarInset>
        </SidebarProvider>
      )}
    </>
  );
}

export default LayoutAdmin;
