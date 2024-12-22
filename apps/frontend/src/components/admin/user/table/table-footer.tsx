import { Button } from '@/components/ui/button';
import PaginationDropdown from '@/components/admin/dropdown-pagination';

const TableFooterUser = ({ table, pagination, setPagination }: any) => {
  const totalPages = Math.ceil(table.getPageCount() / pagination.pageSize);
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
    <div className="flex justify-between items-center w-full">
      <div className="text-sm text-gray-500">
        Showing {pagination.pageIndex * pagination.pageSize + 1} to{' '}
        {Math.min(
          (pagination.pageIndex + 1) * pagination.pageSize,
          table.getPageCount()
        )}{' '}
        of {table.getPageCount()} entries
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
              variant={page === pagination.pageIndex ? 'default' : 'outline'}
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
  );
};

export default TableFooterUser;
