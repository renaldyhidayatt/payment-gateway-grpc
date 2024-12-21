import { useState } from 'react';
import {
  ColumnFiltersState,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useReactTable,
  VisibilityState,
} from '@tanstack/react-table';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { ChevronDown } from 'lucide-react';
import { TableFooterSaldo } from '@/components/admin/saldo/table';
import { TableBodySaldo } from '@/components/admin/saldo/table';
import { saldoColumns } from '@/components/admin/saldo/table/table-column';
import { Table } from '@/components/ui/table';
import { TableHeaderSaldo } from '@/components/admin/saldo/table';
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { saldos } from '@/helpers/saldo_data';

export default function SaldoPage() {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});

  const [pagination, setPagination] = useState({ pageIndex: 0, pageSize: 10 });

  const table = useReactTable({
    data: saldos,
    columns: saldoColumns,
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

  return (
    <div className="flex h-full overflow-hidden">
      <main className="flex-1 p-6 overflow-auto pb-20">
        <div className="flex-1 flex flex-col min-h-0">
          <Card className="w-full shadow-lg rounded-md border">
            <CardHeader className="p-4">
              <div className="flex justify-between items-center">
                <h3 className="text-lg font-semibold">Saldo Table</h3>
                {/* <div className="space-x-2">
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
                </div> */}
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
                      .map((column) => (
                        <DropdownMenuCheckboxItem
                          key={column.id}
                          className="capitalize"
                          checked={column.getIsVisible()}
                          onCheckedChange={(value: any) =>
                            column.toggleVisibility(!!value)
                          }
                        >
                          {column.id}
                        </DropdownMenuCheckboxItem>
                      ))}
                  </DropdownMenuContent>
                </DropdownMenu>
              </div>
              <div className="rounded-md border h-[525px] overflow-y-scroll">
                <Table>
                  <TableHeaderSaldo table={table} />
                  <TableBodySaldo table={table} />
                </Table>
              </div>
            </CardContent>
            <CardFooter className="px-4 py-4 border-t">
              <TableFooterSaldo
                table={table}
                pagination={pagination}
                setPagination={setPagination}
              />
            </CardFooter>
          </Card>
        </div>
      </main>
    </div>
  );
}
