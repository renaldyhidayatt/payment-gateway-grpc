import { useToast } from "@/hooks/use-toast";
// import { handleUnauthorized } from "@/store/auth";

export const handleApiError = (error:  any, setError: any, setLoading: any) => {
//   const { toast } = useToast();
//   const isNetworkError = !error.response;

//   if (isNetworkError) {
//     toast({
//       title: "Aplikasi Sedang Dalam Perawatan",
//       description:
//         "Kami mengalami masalah koneksi dengan server atau server tidak dapat dijangkau. Silakan coba lagi nanti.",
//       variant: "destructive",
//     });

//     if (
//       window.location.pathname !== "/auth/login" &&
//       window.location.pathname !== "/auth/register"
//     ) {
//       // window.location.href = "/auth/login";
//     }

//     setError("Aplikasi sedang dalam perawatan.");
//     setLoading(false);
//     return;
//   }

//   const errorMessage = error.response?.data?.message;
//   const status = error.response?.status;

//   if (status === 401) {
//     // handleUnauthorized(errorMessage);

//     toast({
//       title: "Unauthorized",
//       description: "Anda tidak memiliki izin untuk melakukan tindakan ini.",
//       variant: "destructive",
//     });
//   } else if (status === 422) {
//     const validationErrors = error.response?.data.errors || {};
//     const friendlyErrorMessages = Object.values(validationErrors)
//       .map((msg) => `• ${msg}`)
//       .join("\n"); // Gunakan line break untuk readability

//     toast({
//       title: "Kesalahan Validasi",
//       description: (
//         <>
//           <div>Mohon periksa bagian berikut:</div>
//           <pre style={{ whiteSpace: "pre-wrap", fontSize: "0.9rem" }}>
//             {friendlyErrorMessages}
//           </pre>
//           <div>Ada beberapa kesalahan pada input Anda. Mohon periksa dan coba lagi.</div>
//         </>
//       ),
//       variant: "destructive",
//     });
//   } else if (status === 403) {
//     toast({
//       title: "Forbidden",
//       description: "Anda tidak memiliki izin untuk melakukan tindakan ini.",
//       variant: "destructive",
//     });

//     if (
//       window.location.pathname !== "/auth/login" &&
//       window.location.pathname !== "/auth/register"
//     ) {
      
//     }
//   } else if (status === 404) {
//     toast({
//       title: "Resource Not Found",
//       description: "Sumber daya yang Anda minta tidak ditemukan.",
//       variant: "destructive",
//     });

//     setError("Sumber daya yang diminta tidak ditemukan.");
//   } else if (status === 429) {
//     toast({
//       title: "Too Many Requests",
//       description:
//         "Terlalu banyak permintaan dalam waktu singkat. Silakan coba lagi nanti.",
//       variant: "destructive",
//     });
//   } else {
//     toast({
//       title: "Error",
//       description:
//         errorMessage ||
//         "Terjadi kesalahan yang tidak terduga. Silakan coba lagi nanti.",
//       variant: "destructive",
//     });
//   }

//   setError(errorMessage);
//   setLoading(false);
};
