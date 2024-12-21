import { DropdownMenu } from '@radix-ui/react-dropdown-menu';

import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '../ui/dropdown-menu';

const PaginationDropdown = ({ pagination, setPagination, table }: any) => {
  return (
    <div className="flex items-center space-x-2">
      <span className="text-sm">Rows per page:</span>

      <DropdownMenu>
        <DropdownMenuTrigger className="border rounded px-2 py-1 text-sm">
          {pagination.pageSize} {/* Display the current page size */}
        </DropdownMenuTrigger>

        <DropdownMenuContent className="p-2">
          <DropdownMenuItem
            className="px-4 py-2 hover:bg-gray-200"
            onClick={() => {
              setPagination({ ...pagination, pageSize: 5 });
              table.setPageSize(5);
            }}
          >
            5
          </DropdownMenuItem>
          <DropdownMenuItem
            className="px-4 py-2 hover:bg-gray-200"
            onClick={() => {
              setPagination({ ...pagination, pageSize: 10 });
              table.setPageSize(10);
            }}
          >
            10
          </DropdownMenuItem>
          <DropdownMenuItem
            className="px-4 py-2 hover:bg-gray-200"
            onClick={() => {
              setPagination({ ...pagination, pageSize: 20 });
              table.setPageSize(20);
            }}
          >
            20
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

export default PaginationDropdown;
