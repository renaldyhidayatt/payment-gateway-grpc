import { useEffect, useState } from "react";
import {
  ColumnFiltersState,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useReactTable,
  VisibilityState,
} from "@tanstack/react-table";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ChevronDown } from "lucide-react";
import { Table } from "@/components/ui/table";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  cardColumns,
  TableBodyCard,
  TableFooterCard,
  TableHeaderCard,
} from "@/components/admin/card/table";
import { AddCard } from "@/components/admin/card/modal/CreateModal";
import useModalCard from "@/store/card/modal";
import useCardStore from "@/store/card/card";
import { Skeleton } from "@/components/ui/skeleton";

export default function CardPage() {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const [search, setSearch] = useState("");
  const [isLoadingWithDelay, setIsLoadingWithDelay] = useState(false);
  const [pageSize, setPageSize] = useState(10);
  const [currentPage, setCurrentPage] = useState(1);

  const { showModal } = useModalCard();

  const {
    cards,
    pagination,
    loadingGetCards,
    setLoadingGetCards,
    findAllCards,
  } = useCardStore();

  const table = useReactTable({
    data: cards || [],
    columns: cardColumns,
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
      pagination: {
        pageIndex: currentPage - 1,
        pageSize: pageSize,
      },
    },
    pageCount: pagination.totalPages || 1,
    manualPagination: true,
  });

  useEffect(() => {
    setIsLoadingWithDelay(true);

    const delayTimer = setTimeout(() => {
      setIsLoadingWithDelay(false);
    }, 2000);

    return () => clearTimeout(delayTimer);
  }, [loadingGetCards]);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        setLoadingGetCards(true);
        await findAllCards(search, currentPage, pageSize);
      } catch (error) {
        console.error("Error fetching users:", error);
      } finally {
        setLoadingGetCards(false);
      }
    };

    fetchUsers();
  }, [search, currentPage, pageSize]);

  const handlePageChange = (newPage: number) => {
    console.log("Changing to page:", newPage);
    setCurrentPage(newPage);
  };

  const handlePageSizeChange = (newSize: number) => {
    console.log("Changing page size to:", newSize);
    setPageSize(newSize);
    setCurrentPage(1);
  };

  return (
    <div className="flex h-full overflow-hidden">
      <main className="flex-1 p-6 overflow-auto pb-20">
        <div className="flex-1 flex flex-col min-h-0">
          <Card className="w-full shadow-lg rounded-md border">
            <CardHeader className="p-4">
              <div className="flex justify-between items-center">
                <h3 className="text-lg font-semibold">Table Card</h3>
                <div className="space-x-2">
                  {/* <ImportExcelDialog onImport={0} />
                  <Button
                    // onClick={handleExportExcel}
                    variant="outline"
                    size="sm"
                  >
                    <FileDown className="mr-2 h-4 w-4" />
                    Export Excel
                  </Button> */}

                  <AddCard />
                </div>
              </div>
            </CardHeader>
            <CardContent className="p-4">
              <div className="flex items-center py-4">
                <Input
                  placeholder="Filter by card number"
                  value={search}
                  onChange={(event) => setSearch(event.target.value)}
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
                {isLoadingWithDelay || loadingGetCards ? (
                  <div className="space-y-4">
                    <Skeleton className="h-10 w-full" />
                    <Skeleton className="h-10 w-full" />
                    <Skeleton className="h-10 w-full" />
                    <Skeleton className="h-10 w-full" />
                  </div>
                ) : (
                  <div className="rounded-md border h-[525px] overflow-y-scroll">
                    <Table>
                      <TableHeaderCard table={table} />
                      <TableBodyCard table={table} />
                    </Table>
                  </div>
                )}
              </div>
            </CardContent>
            <CardFooter className="px-4 py-4 border-t">
              <TableFooterCard
                table={table}
                pagination={{
                  currentPage: currentPage,
                  pageSize: pageSize,
                  totalItems: pagination.totalItems,
                  totalPages: pagination.totalPages,
                }}
                onPageChange={handlePageChange}
                onPageSizeChange={handlePageSizeChange}
              />
            </CardFooter>
          </Card>
        </div>
      </main>
    </div>
  );
}
