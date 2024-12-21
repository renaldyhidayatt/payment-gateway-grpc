import { useState } from 'react';
import { Paymentdata } from '@/helpers/payment_data';
import { Checkbox } from '@/components/ui/checkbox';
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from '@tanstack/react-table';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from '@/components/ui/card';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';

import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import {
  ArrowUpDown,
  ChevronDown,
  Eye,
  FileDown,
  MoreHorizontal,
  Pencil,
  Trash,
} from 'lucide-react';
import { Payment } from '@/types/payment';
import { ImportExcelDialog } from '@/components/admin/modal/importExcel';
import { AddEmployee } from '@/components/admin/modal/addEmployee';
import PaginationDropdown from '@/components/admin/dropdown-pagination';

export const columns: ColumnDef<Payment>[] = [
  {
    id: 'select',
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && 'indeterminate')
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: 'status',
    header: 'Status',
    cell: ({ row }) => (
      <div className="capitalize">{row.getValue('status')}</div>
    ),
  },
  {
    accessorKey: 'email',
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === 'asc')}
        >
          Email
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => <div className="lowercase">{row.getValue('email')}</div>,
  },
  {
    accessorKey: 'amount',
    header: () => <div className="text-right">Amount</div>,
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue('amount'));

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
      }).format(amount);

      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: 'actions',
    header: () => <div className="text-right">Actions</div>,
    enableSorting: false,
    enableHiding: false,
    cell: ({ row }) => {
      const payment = row.original;

      return (
        <div className="flex justify-end">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon" className="h-8 w-8">
                <span className="sr-only">Open menu</span>
                <MoreHorizontal className="h-5 w-5" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-48">
              <DropdownMenuLabel>Actions</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => console.log('Viewing details for:', payment.id)}
              >
                <Eye className="mr-2 h-4 w-4 text-gray-500" />
                View Details
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={() => console.log('Editing payment:', payment.id)}
              >
                <Pencil className="mr-2 h-4 w-4 text-gray-500" />
                Edit
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={() => console.log('Deleting payment:', payment.id)}
                className="text-red-600"
              >
                <Trash className="mr-2 h-4 w-4 text-red-500" />
                Delete
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      );
    },
  },
];

export default function TablePayment() {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});

  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });

  const table = useReactTable({
    data: Paymentdata,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
      pagination,
    },
    onPaginationChange: setPagination,
  });

  const totalPages = Math.ceil(Paymentdata.length / pagination.pageSize); // Total pages
  const pageWindowSize = 5;

  const startPage =
    Math.floor(pagination.pageIndex / pageWindowSize) * pageWindowSize;
  const endPage = Math.min(startPage + pageWindowSize, totalPages);

  const handleNextPage = () => {
    if (pagination.pageIndex < totalPages - 1) {
      table.nextPage();
    }
  };

  const handlePreviousPage = () => {
    if (pagination.pageIndex > 0) {
      table.previousPage();
    }
  };

  return (
    <div className="flex h-full overflow-hidden">
      <main className="flex-1 p-6 overflow-auto pb-20">
        <div className="flex-1 flex flex-col min-h-0">
          <Card className="w-full shadow-lg rounded-md border">
            <CardHeader className="p-4">
              <div className="flex justify-between items-center">
                <h3 className="text-lg font-semibold">Table Payment</h3>
                <div className="space-x-2">
                  <ImportExcelDialog onImport={0} />
                  <Button
                    // onClick={handleExportExcel}
                    variant="outline"
                    size="sm"
                  >
                    <FileDown className="mr-2 h-4 w-4" />
                    Export Excel
                  </Button>

                  <AddEmployee onSubmit={0} />
                </div>
              </div>
            </CardHeader>
            <CardContent className="p-4">
              <div className="flex items-center py-4">
                <Input
                  placeholder="Filter emails..."
                  value={
                    (table.getColumn('email')?.getFilterValue() as string) ?? ''
                  }
                  onChange={(event) =>
                    table.getColumn('email')?.setFilterValue(event.target.value)
                  }
                  className="max-w-sm"
                />
                <DropdownMenu>
                  <DropdownMenuTrigger asChild>
                    <Button variant="outline" className="ml-auto">
                      Columns <ChevronDown className="ml-2 h-4 w-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    {table
                      .getAllColumns()
                      .filter((column) => column.getCanHide())
                      .map((column) => {
                        return (
                          <DropdownMenuCheckboxItem
                            key={column.id}
                            className="capitalize"
                            checked={column.getIsVisible()}
                            onCheckedChange={(value) =>
                              column.toggleVisibility(!!value)
                            }
                          >
                            {column.id}
                          </DropdownMenuCheckboxItem>
                        );
                      })}
                  </DropdownMenuContent>
                </DropdownMenu>
              </div>

              <div className="rounded-md border h-[525px] overflow-y-scroll">
                <Table>
                  <TableHeader>
                    {table.getHeaderGroups().map((headerGroup) => (
                      <TableRow key={headerGroup.id}>
                        {headerGroup.headers.map((header) => {
                          return (
                            <TableHead key={header.id}>
                              {header.isPlaceholder
                                ? null
                                : flexRender(
                                    header.column.columnDef.header,
                                    header.getContext()
                                  )}
                            </TableHead>
                          );
                        })}
                      </TableRow>
                    ))}
                  </TableHeader>
                  <TableBody>
                    {table.getRowModel().rows?.length ? (
                      table.getRowModel().rows.map((row) => (
                        <TableRow
                          key={row.id}
                          data-state={row.getIsSelected() && 'selected'}
                        >
                          {row.getVisibleCells().map((cell) => (
                            <TableCell key={cell.id}>
                              {flexRender(
                                cell.column.columnDef.cell,
                                cell.getContext()
                              )}
                            </TableCell>
                          ))}
                        </TableRow>
                      ))
                    ) : (
                      <TableRow>
                        <TableCell
                          colSpan={columns.length}
                          className="h-24 text-center"
                        >
                          No results.
                        </TableCell>
                      </TableRow>
                    )}
                  </TableBody>
                </Table>
              </div>
            </CardContent>

            <CardFooter className="px-4 py-4 border-t">
              <div className="flex justify-between items-center w-full">
                <div className="text-sm text-gray-500">
                  Showing {pagination.pageIndex * pagination.pageSize + 1} to{' '}
                  {Math.min(
                    (pagination.pageIndex + 1) * pagination.pageSize,
                    Paymentdata.length
                  )}{' '}
                  of {Paymentdata.length} entries
                </div>

                <PaginationDropdown
                  pagination={pagination}
                  setPagination={setPagination}
                  table={table}
                />

                <div className="flex items-center space-x-2">
                  <Button
                    variant="outline"
                    onClick={handlePreviousPage}
                    disabled={pagination.pageIndex === 0}
                  >
                    Previous
                  </Button>

                  {[...Array(endPage - startPage).keys()].map((_, i) => {
                    const page = startPage + i;
                    return (
                      <Button
                        key={page}
                        variant={
                          page === pagination.pageIndex ? 'default' : 'outline'
                        }
                        onClick={() => table.setPageIndex(page)}
                      >
                        {page + 1}
                      </Button>
                    );
                  })}

                  {endPage < totalPages && <span>...</span>}

                  <Button
                    variant="outline"
                    onClick={handleNextPage}
                    disabled={pagination.pageIndex === totalPages - 1}
                  >
                    Next
                  </Button>
                </div>
              </div>
            </CardFooter>
          </Card>
        </div>
      </main>
    </div>
  );
}
