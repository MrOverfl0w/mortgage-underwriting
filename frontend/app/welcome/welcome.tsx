import React, { useEffect, useState, useCallback } from "react";

type DecisionRecord = {
  name: string;
  monthly_income: number;
  monthly_debts: number;
  loan_amount: number;
  property_value?: number;
  credit_score: number;
  occupancy: string;
  decision: string;
  dti: number;
  ltv: number;
  reason?: string;
};

export function Welcome() {
  const [records, setRecords] = useState<DecisionRecord[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchRecords = useCallback(() => {
    setLoading(true);
    fetch("/api/loan-history")
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch decisions");
        return res.json();
      })
      .then(setRecords)
      .catch(() => setError("Could not load decisions"))
      .finally(() => setLoading(false));
  }, []);

  useEffect(() => {
    fetchRecords();
  }, [fetchRecords]);

  return (
    <main className="flex items-center justify-center pt-16 pb-4">
      <div className="flex-1 flex flex-col items-center gap-8 min-h-0">
        <header className="flex-col items-center">
          <h1 className="text-2xl/7 font-bold text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">Mortgage Decision's System</h1>
        </header>
        <DecisionMaker onSuccess={fetchRecords} />
        <DecisionList
          records={records}
          loading={loading}
          error={error}
          onRefresh={fetchRecords}
        />
      </div>
    </main>
  );
}

type DecisionMakerProps = {
  onSuccess: () => void;
};

export function DecisionMaker({ onSuccess }: DecisionMakerProps) {
  const [form, setForm] = useState({
    name: "",
    monthly_income: "",
    monthly_debts: "",
    property_value: "",
    loan_amount: "",
    credit_score: "",
    occupancy: "",
  });
  const [result, setResult] = useState<{
    decision: string;
    dti: number;
    ltv: number;
    reason?: string;
  } | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setResult(null);
    try {
      const res = await fetch("/api/request-loan", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name: form.name,
          monthly_income: parseFloat(form.monthly_income),
          monthly_debts: parseFloat(form.monthly_debts),
          property_value: parseFloat(form.property_value),
          loan_amount: parseFloat(form.loan_amount),
          credit_score: parseInt(form.credit_score, 10),
          occupancy: form.occupancy,
        }),
      });
      if (!res.ok) throw new Error("API error");
      const data = await res.json();
      setResult(data);
      onSuccess(); // Refresh the decision list after successful submission
    } catch (err: any) {
      setError("Failed to get decision.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className="max-w-5xl w-full mx-auto mt-8 p-6 border rounded-xl shadow">
      <h2 className="text-xl font-bold mb-4">Calculate and pre-approve your mortgage</h2>
      <form onSubmit={handleSubmit} className="grid grid-cols-1 md:grid-cols-9 gap-x-6 gap-y-6">
        <div className="md:col-span-3">
          <label htmlFor="name" className="block text-sm/6 font-medium text-gray-900">Name</label>
          <input
            className="w-full border p-2 rounded"
            id="name"
            name="name"
            placeholder="Borrower name"
            value={form.name}
            onChange={handleChange}
            required
          />
        </div>
        <div className="md:col-span-2">
          <label htmlFor="monthly_income" className="sblock text-sm/6 font-medium text-gray-900">Monthly Income</label>
          <input
            className="w-full border p-2 rounded"
            id="monthly_income"
            name="monthly_income"
            placeholder="Monthly Income"
            type="number"
            value={form.monthly_income}
            onChange={handleChange}
            required
          />
        </div>
        <div className="md:col-span-2">
          <label htmlFor="monthly_debts" className="block text-sm/6 font-medium text-gray-900">Monthly Debts</label>
          <input
            className="w-full border p-2 rounded"
            id="monthly_debts"
            name="monthly_debts"
            placeholder="Monthly Debts"
            type="number"
            value={form.monthly_debts}
            onChange={handleChange}
            required
          />
        </div>
        <div className="md:col-span-2">
          <label htmlFor="property_value" className="block text-sm/6 font-medium text-gray-900">Property Value</label>
          <input
            className="w-full border p-2 rounded"
            id="property_value"
            name="property_value"
            placeholder="Property Value"
            type="number"
            value={form.property_value}
            onChange={handleChange}
            required
          />
        </div>
        <div className="md:col-span-2">
          <label htmlFor="loan_amount" className="block text-sm/6 font-medium text-gray-900">Loan Amount</label>
          <input
            className="w-full border p-2 rounded"
            id="loan_amount"
            name="loan_amount"
            placeholder="Loan Amount"
            type="number"
            value={form.loan_amount}
            onChange={handleChange}
            required
          />
        </div>
        <div className="md:col-span-2">
          <label htmlFor="credit_score" className="block text-sm/6 font-medium text-gray-900">Credit Score</label>
          <input
            className="w-full border p-2 rounded"
            id="credit_score"
            name="credit_score"
            placeholder="Credit Score"
            type="number"
            value={form.credit_score}
            onChange={handleChange}
            required
          />
        </div>
        <div className="md:col-span-2">
          <label htmlFor="occupancy" className="block text-sm/6 font-medium text-gray-900">Occupancy</label>
          <select
            className="w-full border p-2 rounded"
            id="occupancy"
            name="occupancy"
            value={form.occupancy}
            onChange={handleChange}
            required
          >
            <option value="">Select Occupancy</option>
            <option value="primary">Primary</option>
            <option value="secondary">Secondary</option>
            <option value="investment">Investment</option>
          </select>
        </div>
        <button
          id="submit"
          type="submit"
          className="max-h-max mt-auto rounded-md bg-cyan-600 text-white py-2 rounded hover:bg-cyan-800 col-span-1"
          disabled={loading}
        >
          {loading ? "Checking..." : "Review"}
        </button>
      </form>
      {error && <div className="text-red-600 mt-4">{error}</div>}
      {result && (
        <div className="mt-6 p-4 border rounded bg-gray-50">
          <div>
            <strong>Decision:</strong> {result.decision}
          </div>
          <div>
            <strong>DTI:</strong> {result.dti}
          </div>
          <div>
            <strong>LTV:</strong> {result.ltv}
          </div>
          <div>
            <strong>Reason:</strong> {result.reason}
          </div>
        </div>
      )}
    </section>
  );
}

type DecisionListProps = {
  records: DecisionRecord[];
  loading: boolean;
  error: string | null;
  onRefresh: () => void;
};

export function DecisionList({ records, loading, error }: DecisionListProps) {
  const [expanded, setExpanded] = useState<number | null>(null);

  if (loading) return <div className="mt-8">Loading decisions...</div>;
  if (error) return <div className="mt-8 text-red-600">{error}</div>;
  if (!records.length) return <div className="mt-8">No decisions found.</div>;

  return (
    <section className="max-w-5xl w-full mx-auto mt-12">
      <h2 className="text-lg font-bold mb-4">Previous Decisions</h2>
      <div className="border rounded-xl overflow-hidden">
        {records.map((rec, idx) => (
          <div key={idx} className="border-b last:border-b-0">
            <button
              className="w-full flex items-center justify-between px-4 py-3 hover:bg-gray-100 focus:outline-none"
              onClick={() => setExpanded(expanded === idx ? null : idx)}
              aria-expanded={expanded === idx}
            >
              <span className="flex-1 text-left">
                <span className="font-semibold">{rec.name}</span>
                <span className="ml-4">DTI: <span className="font-mono">{rec.dti.toFixed(2)}</span></span>
                <span className="ml-4">LTV: <span className="font-mono">{rec.ltv.toFixed(2)}</span></span>
                <span className="ml-4 text-gray-600">{rec.reason}</span>
              </span>
              <span className="ml-4">{expanded === idx ? "▲" : "▼"}</span>
            </button>
            {expanded === idx && (
              <div className="bg-gray-50 px-6 py-4 text-sm grid grid-cols-2 gap-x-8 gap-y-2">
                <div><strong>Name:</strong> {rec.name}</div>
                <div><strong>Monthly Income:</strong> {rec.monthly_income}</div>
                <div><strong>Monthly Debts:</strong> {rec.monthly_debts}</div>
                <div><strong>Loan Amount:</strong> {rec.loan_amount}</div>
                <div><strong>Property Value:</strong> {rec.property_value}</div>
                <div><strong>Credit Score:</strong> {rec.credit_score}</div>
                <div><strong>Occupancy:</strong> {rec.occupancy}</div>
                <div><strong>Decision:</strong> {rec.decision}</div>
                <div><strong>DTI:</strong> {rec.dti}</div>
                <div><strong>LTV:</strong> {rec.ltv}</div>
                <div className="col-span-2"><strong>Reason:</strong> {rec.reason}</div>
              </div>
            )}
          </div>
        ))}
      </div>
    </section>
  );
}



const resources = [
  {
    href: "https://reactrouter.com/docs",
    text: "React Router Docs",
    icon: (
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
        className="stroke-gray-600 group-hover:stroke-current dark:stroke-gray-300"
      >
        <path
          d="M9.99981 10.0751V9.99992M17.4688 17.4688C15.889 19.0485 11.2645 16.9853 7.13958 12.8604C3.01467 8.73546 0.951405 4.11091 2.53116 2.53116C4.11091 0.951405 8.73546 3.01467 12.8604 7.13958C16.9853 11.2645 19.0485 15.889 17.4688 17.4688ZM2.53132 17.4688C0.951566 15.8891 3.01483 11.2645 7.13974 7.13963C11.2647 3.01471 15.8892 0.951453 17.469 2.53121C19.0487 4.11096 16.9854 8.73551 12.8605 12.8604C8.73562 16.9853 4.11107 19.0486 2.53132 17.4688Z"
          strokeWidth="1.5"
          strokeLinecap="round"
        />
      </svg>
    ),
  },
  {
    href: "https://rmx.as/discord",
    text: "Join Discord",
    icon: (
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="20"
        viewBox="0 0 24 20"
        fill="none"
        className="stroke-gray-600 group-hover:stroke-current dark:stroke-gray-300"
      >
        <path
          d="M15.0686 1.25995L14.5477 1.17423L14.2913 1.63578C14.1754 1.84439 14.0545 2.08275 13.9422 2.31963C12.6461 2.16488 11.3406 2.16505 10.0445 2.32014C9.92822 2.08178 9.80478 1.84975 9.67412 1.62413L9.41449 1.17584L8.90333 1.25995C7.33547 1.51794 5.80717 1.99419 4.37748 2.66939L4.19 2.75793L4.07461 2.93019C1.23864 7.16437 0.46302 11.3053 0.838165 15.3924L0.868838 15.7266L1.13844 15.9264C2.81818 17.1714 4.68053 18.1233 6.68582 18.719L7.18892 18.8684L7.50166 18.4469C7.96179 17.8268 8.36504 17.1824 8.709 16.4944L8.71099 16.4904C10.8645 17.0471 13.128 17.0485 15.2821 16.4947C15.6261 17.1826 16.0293 17.8269 16.4892 18.4469L16.805 18.8725L17.3116 18.717C19.3056 18.105 21.1876 17.1751 22.8559 15.9238L23.1224 15.724L23.1528 15.3923C23.5873 10.6524 22.3579 6.53306 19.8947 2.90714L19.7759 2.73227L19.5833 2.64518C18.1437 1.99439 16.6386 1.51826 15.0686 1.25995ZM16.6074 10.7755L16.6074 10.7756C16.5934 11.6409 16.0212 12.1444 15.4783 12.1444C14.9297 12.1444 14.3493 11.6173 14.3493 10.7877C14.3493 9.94885 14.9378 9.41192 15.4783 9.41192C16.0471 9.41192 16.6209 9.93851 16.6074 10.7755ZM8.49373 12.1444C7.94513 12.1444 7.36471 11.6173 7.36471 10.7877C7.36471 9.94885 7.95323 9.41192 8.49373 9.41192C9.06038 9.41192 9.63892 9.93712 9.6417 10.7815C9.62517 11.6239 9.05462 12.1444 8.49373 12.1444Z"
          strokeWidth="1.5"
        />
      </svg>
    ),
  },
];
