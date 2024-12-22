import { ColumnDef } from '@tanstack/react-table';
import { Checkbox } from '@/components/ui/checkbox';
import TableActionTransaction from './table-action';
import { Transaction } from '@/types/admin/transaction';

export const transactionColumns: ColumnDef<Transaction>[] = [
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
    accessorKey: 'card_number',
    header: 'Card Number',
    cell: ({ row }) => (
      <div className="font-mono">{row.getValue('card_number')}</div>
    ),
  },
  {
    accessorKey: 'amount',
    header: 'Amount',
    cell: ({ row }) => {
      const amount = row.getValue('amount') as number;
      const formatted = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
      }).format(amount);
      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: 'payment_method',
    header: 'Payment Method',
    cell: ({ row }) => {
      const method = row.getValue('payment_method') as string;
      const formatted =
        method
          .replace(/_/g, ' ')
          .replace(/^\w/, (c) => c.toUpperCase()); // Format: `credit_card` -> `Credit card`
      return <div>{formatted}</div>;
    },
  },
  {
    accessorKey: 'merchant_id',
    header: 'Merchant ID',
    cell: ({ row }) => <div>{row.getValue('merchant_id')}</div>,
  },
  {
    accessorKey: 'transaction_time',
    header: 'Transaction Time',
    cell: ({ row }) => {
      const time = row.getValue('transaction_time') as string;
      return <div>{new Date(time).toLocaleString()}</div>;
    },
  },
  {
    accessorKey: 'created_at',
    header: 'Created At',
    cell: ({ row }) => {
      const createdAt = row.getValue('created_at') as string;
      return <div>{new Date(createdAt).toLocaleString()}</div>;
    },
  },
  {
    accessorKey: 'updated_at',
    header: 'Updated At',
    cell: ({ row }) => {
      const updatedAt = row.getValue('updated_at') as string;
      return <div>{new Date(updatedAt).toLocaleString()}</div>;
    },
  },
  {
    accessorKey: 'deleted_at',
    header: 'Deleted At',
    cell: ({ row }) => {
      const deletedAt = row.getValue('deleted_at') as string | null;
      return <div>{deletedAt ? new Date(deletedAt).toLocaleString() : '-'}</div>;
    },
  },
  {
    accessorKey: 'actions',
    header: () => <div className="text-right">Actions</div>,
    enableSorting: false,
    enableHiding: false,
    cell: ({ row }) => <TableActionTransaction transaction={row.original} />,
  },
];
