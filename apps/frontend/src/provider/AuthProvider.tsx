import { useEffect, useState } from "react";
import { Navigate, useLocation } from "react-router-dom";
import useAuthStore from "../store/auth";
import { jwtDecode} from "jwt-decode";
import usePreviousPath from "@/hooks/utils/usePreviousPath";
import { useToast } from "@/hooks/use-toast";

const AuthProvider = ({ children }: any) => {
  const { refreshAccessToken, isAuthenticated, logout, accessToken } = useAuthStore();
  const { pathname } = useLocation();
  const previousPath = usePreviousPath();
  const [requestedLocation, setRequestedLocation] = useState<string | null>(null);
  const { toast } = useToast();

  useEffect(() => {
    if (!accessToken) {
      toast({
        title: "Error",
        description: "You are not logged in",
        variant: "destructive",
      });
      setRequestedLocation(pathname);
      return;
    }
  
    if (isAuthenticated && accessToken) {
      try {
        const decodedToken: { exp?: number } = jwtDecode(accessToken); 
        const expirationTime = decodedToken.exp ? decodedToken.exp * 1000 : 0;
        const currentTime = Date.now();
        const timeRemaining = expirationTime - currentTime;
  
        if (timeRemaining <= 0) {
          logout();
        } else {
          const timeoutId = setTimeout(async () => {
            try {
              await refreshAccessToken(); 
            } catch (error) {
              logout();
            }
          }, Math.min(timeRemaining - 60000, 15 * 60 * 1000)); 
  
          return () => clearTimeout(timeoutId);
        }
      } catch (error) {
        console.error("Error decoding token:", error);
        logout();
      }
    }
  }, [refreshAccessToken, isAuthenticated, accessToken, logout, pathname, toast]);

  useEffect(() => {
    if (!isAuthenticated && pathname !== requestedLocation) {
      setRequestedLocation(pathname); 
    }
  }, [isAuthenticated, pathname, requestedLocation]);

  if (!isAuthenticated) {
    return (
      <Navigate
        to="/auth/login"
        state={{ from: requestedLocation || previousPath }}
      />
    );
  }

  return children;
};

export default AuthProvider;
