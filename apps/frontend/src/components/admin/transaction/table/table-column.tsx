import { ColumnDef } from '@tanstack/react-table';
import { Checkbox } from '@/components/ui/checkbox';
import { Saldo } from '@/types/admin/saldo';
import TableActionSaldo from './table-action';

export const saldoColumns: ColumnDef<Saldo>[] = [
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
    accessorKey: 'total_balance',
    header: 'Total Balance',
    cell: ({ row }) => {
      const totalBalance = row.getValue('total_balance') as number;
      const formatted = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
      }).format(totalBalance);
      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: 'withdraw_amount',
    header: 'Withdraw Amount',
    cell: ({ row }) => {
      const withdrawAmount = row.getValue('withdraw_amount') as number;
      const formatted = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
      }).format(withdrawAmount);
      return <div className="text-right">{formatted}</div>;
    },
  },
  {
    accessorKey: 'withdraw_time',
    header: 'Withdraw Time',
    cell: ({ row }) => {
      const withdrawTime = row.getValue('withdraw_time') as string;
      return <div>{new Date(withdrawTime).toLocaleString()}</div>;
    },
  },
  {
    accessorKey: 'actions',
    header: () => <div className="text-right">Actions</div>,
    enableSorting: false,
    enableHiding: false,
    cell: ({ row }) => <TableActionSaldo saldo={row.original} />,
  },
];
