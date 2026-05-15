import java.util.Stack;

class Main {
	private static int calPoints(String[] operations) {
		int sum = 0;
		Stack<Integer> scoreRecord = new Stack<>();

		for (String val : operations) {
			switch (val) {
				case "+":
					int firstElem = scoreRecord.peek();
					int secondElem = scoreRecord.get(scoreRecord.size() - 2);
					scoreRecord.push(firstElem + secondElem);
					break;
				case "D":
					scoreRecord.push(scoreRecord.peek() * 2);
					break;
				case "C":
					if (!scoreRecord.isEmpty()) {
						scoreRecord.pop();
					}
					break;
				default:
					int elem = Integer.parseInt(val);
					scoreRecord.push(elem);
					break;

			}

		}

		for (int val : scoreRecord) {
			sum += val;
		}

		return sum;
	}

	public static boolean isValid(String s) {
		char[] charArr = s.toCharArray();
		Stack<Character> stack = new Stack<>();
		for (int i = 0; i < charArr.length; i++) {
			switch (charArr[i]) {
				case '(':
				case '{':
				case '[':
					stack.push(charArr[i]);
					break;
				case ')':
					if (!stack.isEmpty() && stack.peek() == '(') {
						stack.pop();
					} else {
						return false;
					}
					break;
				case '}':
					if (!stack.isEmpty() && stack.peek() == '{') {
						stack.pop();
					} else {
						return false;
					}
					break;
				case ']':
					if (!stack.isEmpty() && stack.peek() == '[') {
						stack.pop();
					} else {
						return false;
					}
					break;
			}
		}
		return stack.isEmpty();
	}

	class MinStack {
		private Stack<Integer> stack;
		private Stack<Integer> minStack;

		public MinStack() {
			this.stack = new Stack<>();
			this.minStack = new Stack<>();
		}

		public void push(int val) {
			stack.push(val);
			if (minStack.isEmpty() || val <= minStack.peek()) {
				minStack.push(val);
			}
		}

		public void pop() {
			if (!stack.isEmpty()) {
				int top = stack.pop();
				if (top == minStack.peek()) {
					minStack.pop();
				}
			} else {
				System.out.println("Stack is empty!");
				return;
			}
		}

		public int top() {
			if (!stack.isEmpty()) {
				return stack.peek();
			} else {
				System.out.println("Stack is empty!");
				return -1;
			}
		}

		public int getMin() {
			return minStack.peek();
		}

	}

	public static void main(String[] args) {
		System.out.println(isValid("([{}])"));
	}
}
