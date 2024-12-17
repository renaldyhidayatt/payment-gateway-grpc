import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import { useState, useEffect } from 'react';
import { Moon, Sun } from 'lucide-react';

export function Navbar() {
  const [darkMode, setDarkMode] = useState(false);

  useEffect(() => {
    if (darkMode) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }, [darkMode]);

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
  };

  return (
    <header className="w-full border-b border-blue-200 bg-white">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 items-center justify-between">
          <div className="flex items-center">
            <Link to="/" className="flex-shrink-0">
              <span className="text-2xl font-bold text-blue-600">PayEase</span>
            </Link>
            <nav className="hidden md:ml-6 md:flex md:space-x-8">
              <Link
                to="/solutions"
                className="text-gray-700 hover:text-blue-600 px-3 py-2 text-sm font-medium"
              >
                Solutions
              </Link>
              <Link
                to="/pricing"
                className="text-gray-700 hover:text-blue-600 px-3 py-2 text-sm font-medium"
              >
                Pricing
              </Link>
              <Link
                to="/developers"
                className="text-gray-700 hover:text-blue-600 px-3 py-2 text-sm font-medium"
              >
                Developers
              </Link>
              <Link
                to="/about"
                className="text-gray-700 hover:text-blue-600 px-3 py-2 text-sm font-medium"
              >
                About
              </Link>
            </nav>
          </div>
          <div className="hidden md:flex md:items-center">
            <Button
              onClick={toggleDarkMode}
              variant="ghost"
              className="text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400"
            >
              {darkMode ? (
                <Sun className="h-5 w-5" />
              ) : (
                <Moon className="h-5 w-5" />
              )}
              <span className="sr-only">
                {darkMode ? 'Switch to light mode' : 'Switch to dark mode'}
              </span>
            </Button>
            <Link
              to="/login"
              className="text-gray-700 hover:text-blue-600 px-3 py-2 text-sm font-medium"
            >
              Log in
            </Link>
            <Button className="ml-4 bg-blue-600 text-white hover:bg-blue-700">
              Sign up
            </Button>
          </div>
          <div className="flex items-center md:hidden">
            <Button variant="ghost" className="text-gray-700">
              <span className="sr-only">Open main menu</span>
              <svg
                className="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </Button>
          </div>
        </div>
      </div>
    </header>
  );
}
