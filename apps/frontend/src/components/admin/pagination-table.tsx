import { Button } from '../ui/button';

const PaginationTable = ({
  pagination,
  setPagination,
  table,
  totalPages,
  startPage,
  endPage,
}: any) => {
  const handlePreviousPage = () => {
    if (pagination.pageIndex > 0) {
      table.setPageIndex(pagination.pageIndex - 1);
    }
  };

  const handleNextPage = () => {
    if (pagination.pageIndex < totalPages - 1) {
      table.setPageIndex(pagination.pageIndex + 1);
    }
  };

  return (
    <div className="flex items-center space-x-2">
      {/* Previous Button */}
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

      {/* Ellipsis */}
      {endPage < totalPages && <span>...</span>}

      {/* Next Button */}
      <Button
        variant="outline"
        onClick={handleNextPage}
        disabled={pagination.pageIndex === totalPages - 1}
      >
        Next
      </Button>
    </div>
  );
};

export default PaginationTable;
